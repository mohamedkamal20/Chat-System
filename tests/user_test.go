package tests

import (
	"Chat-System/handlers"
	"Chat-System/models"
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	user := models.User{
		Email:    "test@example.com",
		Password: "password123",
	}

	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/v1/register", bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/register", handlers.Register).Methods("POST")
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestLogin(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	user := models.User{
		Email:    "test@example.com",
		Password: "password123",
	}

	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/login", handlers.Login).Methods("POST")
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
