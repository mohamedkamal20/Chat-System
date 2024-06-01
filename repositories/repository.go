package repositories

import (
	"Chat-System/models"
)

type UserRepo interface {
	CreateUser(user models.User) error
	GetUserByEmail(email string) (models.User, error)
}

type MessageRepo interface {
}