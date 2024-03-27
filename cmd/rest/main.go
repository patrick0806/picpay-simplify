package main

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/patrick0806/picpay-simplify/config/database"
	"github.com/patrick0806/picpay-simplify/internal/controllers"
	"github.com/patrick0806/picpay-simplify/internal/repositories"
	"github.com/patrick0806/picpay-simplify/internal/usecases"
)

func main() {
	db, err := database.OpenConnection()
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8080", loadRoutes(db))
}

func loadRoutes(db *sql.DB) *mux.Router {
	userRepository := repositories.NewUserRepositoryImpl(db)
	transactionRepository := repositories.NewTransactionRepositoryImpl(db)
	userController := controllers.NewUserController(usecases.NewCreateUserUseCase(userRepository))
	transactionController := controllers.NewTransactionController(usecases.NewCreateTransactionUseCase(transactionRepository, userRepository))

	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()

	//Users
	router.HandleFunc("/v1/users", userController.CreateUser).Methods("POST")

	//Transactions
	router.HandleFunc("/v1/transactions", transactionController.CreateTransaction).Methods("POST")
	return r
}
