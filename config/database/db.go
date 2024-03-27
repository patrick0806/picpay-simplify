package database

import "database/sql"

func OpenConnection() (*sql.DB, error) {
	//TODO- change for env vars
	db, err := sql.Open("postgres", "user=root password=123 dbname=picpay sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	return db, err
}
