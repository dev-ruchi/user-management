package main

import (
	"github.com/dev-ruchi/user-management/backend/api"
	"github.com/dev-ruchi/user-management/backend/app"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	app.SetupDatabase()

	api.NewRouter()

	defer app.Db.Close()
}
