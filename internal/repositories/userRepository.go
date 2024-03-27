package repositories

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/patrick0806/picpay-simplify/internal/entities"
)

type UserRepository interface {
	FindByEmailOrDocument(email string, document string) (*entities.User, error)
	FindById(id uuid.UUID) (*entities.User, error)
	Save(user *entities.User) error
}

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

func (u *UserRepositoryImpl) FindByEmailOrDocument(email string, document string) (*entities.User, error) {

	var user entities.User
	err := u.DB.QueryRow("SELECT * FROM users WHERE email = $1 OR document = $2", email, document).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Type, &user.Document, &user.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u *UserRepositoryImpl) FindById(id uuid.UUID) (*entities.User, error) {

	var user entities.User
	err := u.DB.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Type, &user.Document, &user.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserRepositoryImpl) Save(user *entities.User) error {

	_, err := u.DB.Exec("INSERT INTO users (id,name, email, password, type, document) VALUES ($1, $2, $3, $4, $5, $6)",
		user.Id, user.Name, user.Email, user.Password, user.Type, user.Document)
	if err != nil {
		return err
	}

	return nil
}
