package models

import (
	"github.com/gocql/gocql"
	"time"
)

type Message struct {
	MessageID gocql.UUID `json:"message_id,omitempty" bson:"message_id,omitempty"`
	Sender    string     `json:"sender,omitempty" bson:"sender,omitempty"`
	Recipient string     `json:"recipient,omitempty" bson:"recipient,omitempty"`
	Content   string     `json:"content,omitempty" bson:"content,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

func (t Message) MessageResponse() map[string]interface{} {
	return map[string]interface{}{
		"sender":    t.Sender,
		"recipient": t.Recipient,
		"content":   t.Content,
		"timestamp": t.CreatedAt,
	}
}

func (t Message) MessageErrorResponse(error string) map[string]interface{} {
	return map[string]interface{}{
		"errorMessage": error,
	}
}
