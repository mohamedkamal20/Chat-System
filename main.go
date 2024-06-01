package main

import (
	"Chat-System/handlers"
	"Chat-System/repositories/user"
	"Chat-System/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	utils.InitCassandra()
	defer utils.Session.Close()

	userRepo := user.NewUserRepository()
	handlers.InitUserHandlers(userRepo)

	r := mux.NewRouter()

	// Apply middleware for JWT token authentication
	//r.Use(middlewares.AuthMiddleware)

	// User routes
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
