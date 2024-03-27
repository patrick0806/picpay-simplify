package usecases

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/patrick0806/picpay-simplify/internal/entities"
	"github.com/patrick0806/picpay-simplify/internal/repositories"
	"github.com/patrick0806/picpay-simplify/pkg/enums"
	"github.com/patrick0806/picpay-simplify/pkg/utils"
)

type CreateUserUseCase struct {
	UserRepository repositories.UserRepository
}

func NewCreateUserUseCase(repositoy repositories.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: repositoy,
	}
}

func (u *CreateUserUseCase) Execute(user *entities.User) (*entities.User, error) {

	if user.Type != enums.UserType(enums.Shopkeeper).String() && user.Type != enums.UserType(enums.Common).String() {
		return nil, &entities.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid user type",
			Details:    fmt.Sprintf("%v", user.Type),
		}
	}

	if user.Type == enums.UserType(enums.Shopkeeper).String() && !utils.IsValidCNPJ(user.Document) {
		return nil, &entities.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid Document",
			Details:    "Invalid CNPJ.",
		}
	}

	if user.Type == enums.UserType(enums.Common).String() && !utils.IsValidCPF(user.Document) {
		return nil, &entities.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid Document",
			Details:    "Invalid CPF.",
		}
	}
	user.Document = utils.RemoveSpecialCharacters(user.Document)

	findedUser, err := u.UserRepository.FindByEmailOrDocument(user.Email, user.Document)
	if err != nil {
		return nil, &entities.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Error on find user",
			Details:    fmt.Sprintf("%v", err),
		}
	}

	if findedUser != nil {
		return nil, &entities.Error{
			StatusCode: http.StatusConflict,
			Message:    "User already exists",
			Details:    "User with same email or document already exists.",
		}
	}

	user.Id = uuid.New()
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		return nil, &entities.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error on hash password",
			Details:    fmt.Sprintf("%v", err),
		}
	}

	err = u.UserRepository.Save(user)
	if err != nil {
		return nil, &entities.Error{
			StatusCode: http.StatusConflict,
			Message:    "Fail to save user",
			Details:    fmt.Sprintf("%v", user.Type),
		}
	}

	return user, nil
}
