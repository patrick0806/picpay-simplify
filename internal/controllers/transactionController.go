package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/patrick0806/picpay-simplify/internal/entities"
	"github.com/patrick0806/picpay-simplify/internal/usecases"
)

type TransactionController struct {
	CreateTransactionUseCase *usecases.CreateTransactionUseCase
}

func NewTransactionController(createTransactionUseCase *usecases.CreateTransactionUseCase) *TransactionController {
	return &TransactionController{
		CreateTransactionUseCase: createTransactionUseCase,
	}
}

func (t *TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	transaction := entities.Transaction{}

	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entities.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Error on decode body",
			Details:    fmt.Sprintf("%v", err),
		})
		return
	}

	savedTransaction, err := t.CreateTransactionUseCase.Execute(&transaction)

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
	json.NewEncoder(w).Encode(savedTransaction)
}
