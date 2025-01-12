package main

import (
	"log"

	"github.com/dev-ruchi/user-management/backend/api"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	db, err := NewDatabase()

	if err != nil {
		log.Fatalf("Failed to conenct to the database: %v", err)
	}

	router := api.NewRouter(db)
	router.SetupRoutes()

	defer db.Close()
}
