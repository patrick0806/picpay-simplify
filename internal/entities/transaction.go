package entities

import "github.com/google/uuid"

type Transaction struct {
	Id         uuid.UUID `json:"id"`
	PayerId    uuid.UUID `json:"payerId"`
	ReceiverId uuid.UUID `json:"receiverId"`
	Value      float64   `json:"value"`
}
