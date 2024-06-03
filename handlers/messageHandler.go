package handlers

import (
	"Chat-System/models"
	"Chat-System/repositories"
	"Chat-System/services/redis"
	"Chat-System/utils"
	"encoding/json"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var messageRepo repositories.MessageRepo

func InitMessageHandlers(repo repositories.MessageRepo) {
	messageRepo = repo
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	var message models.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check if not authorized
	tokenEmail, ok := r.Context().Value("email").(string)
	if !ok || (tokenEmail != message.Sender) {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	// Check if sender and recipient exist

	// Get sender user from cache
	cachedSenderUser, err := redis.GetUser(message.Sender)
	if cachedSenderUser == nil {
		_, err = userRepo.GetUserByEmail(message.Sender)
		if err != nil {
			http.Error(w, "Sender not found", http.StatusNotFound)
			return
		}
	}
	// Get Recipient user from cache
	cachedRecipientUser, err := redis.GetUser(message.Recipient)
	if cachedRecipientUser == nil {
		_, err = userRepo.GetUserByEmail(message.Recipient)
		if err != nil {
			http.Error(w, "Recipient not found", http.StatusNotFound)
			return
		}
	}

	message.MessageID = gocql.TimeUUID()
	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()
	err = messageRepo.CreateMessage(message)
	if err != nil {
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	// Remove key from cache
	redis.InvalidateCacheMessages(message.Sender)
	redis.InvalidateCacheMessages(message.Recipient)

	w.WriteHeader(http.StatusCreated)
}

func GetMessageHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	// Validate the email parameter format
	if !utils.IsValidEmail(email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Check if not authorized
	tokenEmail, ok := r.Context().Value("email").(string)
	if !ok || (tokenEmail != email) {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	// Check cached messages by email
	messages, err := redis.GetMessages(email)
	if err != nil {
		messages, err = messageRepo.GetMessagesByEmail(email)
		if err != nil {
			http.Error(w, "Failed to retrieve messages", http.StatusInternalServerError)
			return
		}
		redis.SetMessages(email, messages)
	}
	json.NewEncoder(w).Encode(messages)
}
