package handlers

import (
	"Chat-System/models"
	"Chat-System/repositories"
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

	// Check if sender and recipient exist
	_, err = userRepo.GetUserByEmail(message.Sender)
	if err != nil {
		http.Error(w, "Sender not found", http.StatusNotFound)
		return
	}
	_, err = userRepo.GetUserByEmail(message.Recipient)
	if err != nil {
		http.Error(w, "Recipient not found", http.StatusNotFound)
		return
	}

	message.MessageID = gocql.TimeUUID()
	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()
	err = messageRepo.CreateMessage(message)
	if err != nil {
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetMessageHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	messages, err := messageRepo.GetMessagesByEmail(email)
	if err != nil {
		http.Error(w, "Failed to retrieve messages", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(messages)
}
