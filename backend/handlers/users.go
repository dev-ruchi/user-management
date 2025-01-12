package handlers

import (
	"fmt"
	"os"

	"github.com/dev-ruchi/user-management/backend/app"
	"github.com/dev-ruchi/user-management/backend/models"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleLogin(c *fiber.Ctx) error {

	var loginRequest LoginRequest

	if err := c.BodyParser(&loginRequest); err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad requset",
		})
	}

	var user models.User

	query := `
        SELECT id, name, email, password
        FROM users
        WHERE email = $1`

	// Query the database to find the user by username
	row := app.Db.QueryRow(query, loginRequest.Email)

	if row.Err() != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))

	if err != nil {

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	tokenStr, err := generateJWT(user.Id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}

	// If the password is correct, return a success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   tokenStr,
		"user": fiber.Map{
			"id":       user.Id,
			"email":    user.Email,
			"password": user.Password,
		},
	})

}

func HandleAddUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		fmt.Printf("Body parsing error: %v\n", err)
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	query := `
        INSERT INTO users (name, email, password)
        VALUES ($1, $2, $3)
        RETURNING id, name, email, password`

	password, err := (bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost))

	if err != nil {
		fmt.Printf("Password hasing error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong while hashing the password",
		})
	}

	user.Password = string(password)

	err = app.Db.QueryRow(query, user.Name, user.Email, user.Password).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		fmt.Printf("Query error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}
	tokenStr, err := generateJWT(user.Id)

	if err != nil {
		fmt.Printf("JWT generation error: %v\n", err) 
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}


	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Signup successful",
		"token":   tokenStr,
		"user": fiber.Map{
			"id":    user.Id,
			"name":  user.Name,
			"email": user.Email,
			"password": user.Password,
		},
	})

}

func generateJWT(userId int) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"userId":    userId,
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(secret))
	return tokenStr, err
}
