package entities

import "fmt"

type Error struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Details    string `json:"details"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("StatusCode: %d, Message: %s, Details: %s", e.StatusCode, e.Message, e.Details)
}
