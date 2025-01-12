package api

import (
	"log"

	"github.com/dev-ruchi/user-management/backend/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewRouter() {
	

	// Create a new Fiber app
	app := fiber.New()

	// Middleware for CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Authorization, Accept, Origin, Cache-Control, X-Requested-With",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Define routes
	app.Post("/signup", func(c *fiber.Ctx) error {
		return handlers.HandleAddUser(c)
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		return handlers.HandleLogin(c)
	})
	err := app.Listen(":8080")
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
