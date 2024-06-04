package message

import (
	"Chat-System/models"
	"Chat-System/utils"
	"encoding/base64"
	"log"
	"sort"
)

type MessageRepository interface {
	CreateMessage(message models.Message) error
	GetMessagesByEmail(email string, count int, pageState []byte) (MessagesResponse, error)
}

type messageRepository struct{}

func NewMessageRepository() MessageRepository {
	return &messageRepository{}
}

func (r *messageRepository) CreateMessage(message models.Message) error {
	query := "INSERT INTO messages (message_id, sender, recipient, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	if err := utils.Session.Query(query, message.MessageID, message.Sender, message.Recipient, message.Content, message.CreatedAt, message.UpdatedAt).Exec(); err != nil {
		log.Println("Error creating message:", err)
		return err
	}
	return nil
}

func (r *messageRepository) GetMessagesByEmail(email string, count int, pageState []byte) (MessagesResponse, error) {
	var response []map[string]interface{}
	var messages []models.Message

	senderQuery := "SELECT sender, recipient, content, created_at FROM messages WHERE sender = ? LIMIT ? ALLOW FILTERING"
	iter := utils.Session.Query(senderQuery, email, count).PageSize(count).PageState(pageState).Iter()
	scanner := iter.Scanner()
	var message models.Message
	result := MessagesResponse{
		Messages:  nil,
		PageState: "",
	}

	for scanner.Next() {
		err := scanner.Scan(&message.Sender, &message.Recipient, &message.Content, &message.CreatedAt)
		if err != nil {
			log.Println("Error retrieving messages:", err)
			return result, err
		}
		messages = append(messages, message)
	}
	if err := scanner.Err(); err != nil {
		log.Println("Error retrieving messages:", err)
		return result, err
	}

	pageState = iter.PageState()
	var nextPageState string

	nextPageState = base64.StdEncoding.EncodeToString(pageState)

	recipientQuery := "SELECT sender, recipient, content, created_at FROM messages WHERE recipient = ? LIMIT ? ALLOW FILTERING"
	scanner = utils.Session.Query(recipientQuery, email, count).PageSize(count).PageState(pageState).Iter().Scanner()

	for scanner.Next() {
		err := scanner.Scan(&message.Sender, &message.Recipient, &message.Content, &message.CreatedAt)
		if err != nil {
			log.Println("Error retrieving messages:", err)
			return result, err
		}
		messages = append(messages, message)
	}
	if err := scanner.Err(); err != nil {
		log.Println("Error retrieving messages:", err)
		return result, err
	}

	// Sort results by CreatedAt in descending order
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].CreatedAt.After(messages[j].CreatedAt)
	})

	if len(messages) > count {
		messages = messages[:count]
	}

	// Convert messages to response format
	for _, message := range messages {
		response = append(response, message.MessageResponse())
	}

	result = MessagesResponse{
		Messages:  response,
		PageState: nextPageState,
	}

	return result, nil
}

type MessagesResponse struct {
	Messages  []map[string]interface{} `json:"messages"`
	PageState string                   `json:"page_state,omitempty"`
}
