package repositories

import (
	"database/sql"

	"github.com/patrick0806/picpay-simplify/internal/entities"
)

type TransactionRepository interface {
	Save(transaction *entities.Transaction, payerBalance float64, receiverBalance float64) error
}

type TransactionRepositoryImpl struct {
	DB *sql.DB
}

func NewTransactionRepositoryImpl(db *sql.DB) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{DB: db}
}

func (t *TransactionRepositoryImpl) Save(transaction *entities.Transaction, payerBalance float64, reciverBalance float64) error {

	tx, err := t.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO transactions (id, payer_id, reciver_id, value) VALUES ($1, $2, $3, $4)",
		transaction.Id, transaction.PayerId, transaction.ReceiverId, transaction.Value)
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE users SET balance = $1 WHERE id = $2", payerBalance, transaction.PayerId)
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE users SET balance = $1 WHERE id = $2", reciverBalance, transaction.ReceiverId)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
