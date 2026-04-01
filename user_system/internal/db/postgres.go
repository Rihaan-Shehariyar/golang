package db

import (
	"database/sql"
	"fmt"
)

func ConnectDB() (*sql.DB, error) {
	connStr := "postgres://user:Rihaan@123localhost:5432/ecommerce?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected To Postgres")
	return db, nil

}
