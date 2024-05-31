package handlers

import (
	"Chat-System/utils"
	"encoding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

func Login(w http.ResponseWriter, r *http.Request) {
	token, err := utils.GenerateJWT("user@example.com")
	if err != nil {
		http.Error(w, "Failed to generate JWT token", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
	//w.WriteHeader(http.StatusOK)
}
