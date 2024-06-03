package tests

import (
	"Chat-System/handlers"
	"Chat-System/models"
	"Chat-System/utils"
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	message := models.Message{
		Sender:    "test1@example.com",
		Recipient: "test2@example.com",
		Content:   "Hello!",
	}

	jsonMessage, _ := json.Marshal(message)
	req, _ := http.NewRequest("POST", "/api/v1/send", bytes.NewBuffer(jsonMessage))
	req.Header.Set("Content-Type", "application/json")
	token, _ := utils.GenerateJWT(message.Sender)
	bearer := "Bearer " + token
	req.Header.Set("Authorization", bearer)

	res := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/send", handlers.SendMessage).Methods("POST")
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestGetMessages(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	req, _ := http.NewRequest("GET", "/api/v1/messages/test1@example.com", nil)
	req.Header.Set("Content-Type", "application/json")
	token, _ := utils.GenerateJWT("test1@example.com")
	bearer := "Bearer " + token
	req.Header.Set("Authorization", bearer)

	res := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/messages/{email}", handlers.GetMessageHistory).Methods("GET")
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
