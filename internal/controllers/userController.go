package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/patrick0806/picpay-simplify/internal/entities"
	"github.com/patrick0806/picpay-simplify/internal/usecases"
	"github.com/patrick0806/picpay-simplify/pkg/dtos"
)

type UserController struct {
	CreateUserUseCase *usecases.CreateUserUseCase
}

func NewUserController(createUserUseCase *usecases.CreateUserUseCase) *UserController {
	return &UserController{
		CreateUserUseCase: createUserUseCase,
	}
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := dtos.CreateUser{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Error on decode body",
			Details:    fmt.Sprintf("%v", err),
		})
		return
	}

	savedUser, err := u.CreateUserUseCase.Execute(&entities.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Type:     user.Type,
		Document: user.Document,
	})

	if err != nil {
		if customErr, ok := err.(*entities.Error); ok {
			w.WriteHeader(customErr.StatusCode)
			json.NewEncoder(w).Encode(customErr)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entities.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    "Internal server error",
				Details:    fmt.Sprintf("%v", err),
			})
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(savedUser)
}
