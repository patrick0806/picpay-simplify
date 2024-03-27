package dtos

type CreateUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     string `json:"type"`
	Document string `json:"document"`
}
