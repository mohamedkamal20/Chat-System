package models

import "time"

type User struct {
	Email     string    `json:"email,omitempty" validate:"required,email" bson:"email,omitempty"`
	Password  string    `json:"password,omitempty"  validate:"required,min=5" bson:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

func (t User) UserResponse() map[string]interface{} {
	return map[string]interface{}{
		"email":    t.Email,
		"password": t.Password,
	}
}

func (t User) UserErrorResponse(error string) map[string]interface{} {
	return map[string]interface{}{
		"errorMessage": error,
	}
}
