package entities

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	Type     string    `json:"type"`
	Document string    `json:"document"`
	Balance  float64   `json:"balance"`
}
