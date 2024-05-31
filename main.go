package main

import (
	"Chat-System/handlers"
	"Chat-System/middlewares"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	// Apply middleware for JWT token authentication
	r.Use(middlewares.AuthMiddleware)

	// User routes
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
