package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func NewDatabase() (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	err = createUsersTable(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func createUsersTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT,
	        email TEXT,
			password TEXT
		)
	`)

	return err
}
