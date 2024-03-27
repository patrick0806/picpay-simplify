package usecases

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/patrick0806/picpay-simplify/internal/entities"
	"github.com/patrick0806/picpay-simplify/internal/repositories"
)

type CreateTransactionUseCase struct {
	TransactionRepository repositories.TransactionRepository
	UserRepository        repositories.UserRepository
}

func NewCreateTransactionUseCase(transactionsRepository repositories.TransactionRepository, usersRepository repositories.UserRepository) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionRepository: transactionsRepository,
		UserRepository:        usersRepository,
	}
}

func (ct *CreateTransactionUseCase) Execute(transaction *entities.Transaction) (*entities.Transaction, error) {

	if transaction.Value <= 0 {
		return nil, &entities.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid transaction value",
			Details:    fmt.Sprintf("%v", transaction.Value),
		}
	}

	payer, err := ct.UserRepository.FindById(transaction.PayerId)
	if err != nil {
		return nil, &entities.Error{
			StatusCode: http.StatusNotFound,
			Message:    "payer not found",
			Details:    fmt.Sprintf("%v", err),
		}
	}

	if payer.Balance < transaction.Value {
		return nil, &entities.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid transaction value",
			Details:    "Payer does not have enough balance",
		}
	}

	receiver, err := ct.UserRepository.FindById(transaction.ReceiverId)
	if err != nil {
		return nil, &entities.Error{
			StatusCode: http.StatusNotFound,
			Message:    "receiver not found",
			Details:    fmt.Sprintf("%v", err),
		}
	}

	receiver.Balance += transaction.Value
	payer.Balance -= transaction.Value
	transaction.Id = uuid.New()

	err = ct.TransactionRepository.Save(transaction, payer.Balance, receiver.Balance)
	if err != nil {
		return nil, &entities.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    "error on saving transaction",
			Details:    fmt.Sprintf("%v", err),
		}
	}

	return transaction, nil
}
