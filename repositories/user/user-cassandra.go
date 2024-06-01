package user

import (
	"Chat-System/models"
	"Chat-System/repositories"
	"Chat-System/utils"
	"log"
)

type userRepository struct{}

func NewUserRepository() repositories.UserRepo {
	return &userRepository{}
}

func (r *userRepository) CreateUser(user models.User) error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	if err := utils.Session.Query(query, user.Email, user.Password).Exec(); err != nil {
		log.Println("Error creating user:", err)
		return err
	}
	return nil
}

func (r *userRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := "SELECT email, password FROM users WHERE email = ? LIMIT 1"
	if err := utils.Session.Query(query, email).Scan(&user.Email, &user.Password); err != nil {
		log.Println("Error getting user:", err)
		return models.User{}, err
	}
	return user, nil
}
