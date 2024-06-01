package main

import (
	"Chat-System/handlers"
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

	// Apply middleware for JWT token authentication
	//r.Use(middlewares.AuthMiddleware)

	// User routes
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/send", handlers.SendMessage).Methods("POST")
	r.HandleFunc("/message-history/{email}", handlers.GetMessageHistory).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
