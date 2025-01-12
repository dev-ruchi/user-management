package handlers

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/dev-ruchi/user-management/backend/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func HandleLogin(c *fiber.Ctx, db *sql.DB) error {
	var loginRequest LoginRequest

	if err := c.BodyParser(&loginRequest); err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	if err := validate.Struct(loginRequest); err != nil {
		return sendValidationErrors(c, err)
	}

	var user models.User

	query := `
        SELECT id, name, email, password
        FROM users
        WHERE email = $1`

	row := db.QueryRow(query, loginRequest.Email)

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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   tokenStr,
		"user": fiber.Map{
			"id":       user.Id,
			"email":    user.Email,
		},
	})
}

func HandleSignup(c *fiber.Ctx, db *sql.DB) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	if err := validate.Struct(user); err != nil {
		return sendValidationErrors(c, err)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Password hashing error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}

	user.Password = string(password)

	query := `
        INSERT INTO users (name, email, password)
        VALUES ($1, $2, $3)
        RETURNING id, name, email`

	err = db.QueryRow(query, user.Name, user.Email, user.Password).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
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
			"id":       user.Id,
			"name":     user.Name,
			"email":    user.Email,
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

func sendValidationErrors(c *fiber.Ctx, err error) error {
	errs := err.(validator.ValidationErrors)
	errorMessages := make(map[string]string)
	for _, e := range errs {
		errorMessages[e.Field()] = fmt.Sprintf("Field '%s' failed on the '%s' tag", e.Field(), e.Tag())
	}
	return c.Status(fiber.StatusBadRequest).JSON(errorMessages)
}
