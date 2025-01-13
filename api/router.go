package api

import (
	"database/sql"
	"log"
"os"
	"github.com/dev-ruchi/user-management/backend/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	 "github.com/joho/godotenv"
)

type Router struct {
	db *sql.DB
}

func NewRouter(db *sql.DB) *Router {
	return &Router{
		db,
	}
}

func (r *Router) SetupRoutes() {
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
		return handlers.HandleSignup(c, r.db)
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		return handlers.HandleLogin(c, r.db)
	})
	
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	port := os.Getenv("PORT")

    err = app.Listen(port)
    if err != nil {
        log.Fatal("Error starting server: ", err)
    }
}
