package message

import (
	"Chat-System/models"
	"Chat-System/utils"
	"log"
)

type MessageRepository interface {
	CreateMessage(message models.Message) error
	GetMessagesByEmail(email string) ([]map[string]interface{}, error)
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

func (r *messageRepository) GetMessagesByEmail(email string) ([]map[string]interface{}, error) {
	var response []map[string]interface{}
	query := "SELECT sender, recipient, content, created_at FROM messages WHERE sender = ? ALLOW FILTERING"
	scanner := utils.Session.Query(query, email).Iter().Scanner()
	var message models.Message

	for scanner.Next() {
		err := scanner.Scan(&message.Sender, &message.Recipient, &message.Content, &message.CreatedAt)
		if err != nil {
			log.Println("Error retrieving messages:", err)
			return nil, err
		}
		response = append(response, message.MessageResponse())
	}
	if err := scanner.Err(); err != nil {
		log.Println("Error retrieving messages:", err)
		return nil, err
	}
	return response, nil
}
