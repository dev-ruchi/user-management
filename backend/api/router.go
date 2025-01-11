package api

import (
	"log"

	"github.com/dev-ruchi/user-management/backend/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	//"github.com/gin-gonic/gin"
)

func NewRouter() {
	//router := gin.Default()

	// Create a new Fiber app
	app := fiber.New()

	// Middleware for CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Authorization, Accept, Origin, Cache-Control, X-Requested-With",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	//router.MaxMultipartMemory = 8 << 20 // 8 MiB

	//router.Use(corsMiddleware())

	// Define routes
	app.Post("/users", func(c *fiber.Ctx) error {
		return handlers.HandleAddUser(c)
	})
	app.Get("/users", func(c *fiber.Ctx) error {
		return handlers.HandleFetchUsers(c)
	})
	app.Put("/users/:id", func(c *fiber.Ctx) error {
		return handlers.HandleUpdateUsers(c)
	})
	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		return handlers.HandleDeleteUsers(c)
	})
	err := app.Listen(":8080")
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

