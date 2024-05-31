package handlers

import (
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
