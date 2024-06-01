package models

type User struct {
	Email     string `json:"email,omitempty" bson:"email,omitempty"`
	Password  string `json:"password,omitempty" bson:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
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
