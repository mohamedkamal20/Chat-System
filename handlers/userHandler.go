package handlers

import (
	"Chat-System/models"
	"Chat-System/repositories"
	"Chat-System/services/redis"
	"Chat-System/utils"
	"encoding/json"
	"github.com/gocql/gocql"
	"net/http"
	"time"
)

var userRepo repositories.UserRepo

func InitUserHandlers(repo repositories.UserRepo) {
	userRepo = repo
}

func Register(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate user input
	err = utils.ValidateStruct(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get user from cache
	cachedUser, err := redis.GetUser(user.Email)
	if cachedUser == nil {
		// Fetch it from DB
		_, err := userRepo.GetUserByEmail(user.Email)
		if err == nil {
			http.Error(w, "User already exists", http.StatusNotFound)
			return
		}
	} else {
		http.Error(w, "User already exists", http.StatusNotFound)
		return
	}

	user.UserId = gocql.TimeUUID()
	user.Password = utils.HashPasswordMD5(user.Password)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = userRepo.CreateUser(user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Add user to cache
	redis.SetUser(user.Email, user)

	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Get user from cache
	cachedUser, err := redis.GetUser(user.Email)
	if cachedUser == nil {
		// Fetch it from DB
		storedUser, err := userRepo.GetUserByEmail(user.Email)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		if storedUser.Password != utils.HashPasswordMD5(user.Password) {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}
		// Add user to cache
		redis.SetUser(storedUser.Email, storedUser)
	} else {
		if cachedUser.Password != utils.HashPasswordMD5(user.Password) {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		http.Error(w, "Failed to generate JWT token", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
