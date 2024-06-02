package main

import (
	"Chat-System/handlers"
	"Chat-System/middlewares"
	"Chat-System/repositories/message"
	"Chat-System/repositories/user"
	"Chat-System/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	utils.InitCassandra()
	defer utils.Session.Close()

	utils.InitValidator()

	userRepo := user.NewUserRepository()
	handlers.InitUserHandlers(userRepo)

	messageRepo := message.NewMessageRepository()
	handlers.InitMessageHandlers(messageRepo)

	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/api/v1/register", handlers.Register).Methods("POST")
	r.HandleFunc("/api/v1/login", handlers.Login).Methods("POST")

	protectedRoutes := r.PathPrefix("/").Subrouter()
	// Apply middleware for JWT token authentication
	protectedRoutes.Use(middlewares.AuthMiddleware)

	// Message routes
	protectedRoutes.HandleFunc("/api/v1/send", handlers.SendMessage).Methods("POST")
	protectedRoutes.HandleFunc("/api/v1/messages/{email}", handlers.GetMessageHistory).Methods("GET")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
