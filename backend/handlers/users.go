package handlers

import (
	"fmt"
	"log"

	"github.com/dev-ruchi/user-management/backend/app"
	"github.com/dev-ruchi/user-management/backend/models"

	"github.com/gin-gonic/gin"
)

func HandleAddUser(context *gin.Context) {
	var user models.User

	err := context.BindJSON(&user)

	if err != nil {
		fmt.Println(err)
		context.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	query := `
        INSERT INTO users (name, email, date_of_birth)
        VALUES ($1, $2, $3)
        RETURNING id, name, email, date_of_birth`

	err = app.Db.QueryRow(query, user.Name, user.Email, user.DateOfBirth).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.DateOfBirth,
	)

	if err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	context.JSON(201, user)

}

func HandleFetchUsers(context *gin.Context) {
	rows, err := app.Db.Query("SELECT * FROM users")

	if err != nil {

		log.Fatal(err)

		context.JSON(500, gin.H{
			"message": "Something went wrong",
		})
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.DateOfBirth); err != nil {

			log.Fatal(err)

			context.JSON(500, gin.H{
				"message": "Something went wrong",
			})
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {

		log.Fatal(err)

		context.JSON(500, gin.H{
			"message": "Something went wrong",
		})
	}

	context.JSON(200, users)

}

func HandleUpdateUsers(context *gin.Context) {
	// Get the user ID from URL parameters
	id := context.Param("id")

	var user models.User

	// Bind JSON data to the user model
	if err := context.BindJSON(&user); err != nil {
		context.JSON(400, gin.H{
			"message": "Bad request",
		})
		return
	}

	// Update query
	query := `
        UPDATE users 
        SET name=$1, email=$2, date_of_birth=$3
        WHERE id=$4
        RETURNING id, name, email, date_of_birth`

	// Execute the query and update the database
	err := app.Db.QueryRow(query, user.Name, user.Email, user.DateOfBirth, id).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.DateOfBirth,
	)

	if err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	// Respond with the updated user details
	context.JSON(200, user)
}

func HandleDeleteUsers(context *gin.Context) {

	query := `
      DELETE FROM users WHERE id=$1;`

	_, err := app.Db.Query(query, context.Param("id"))

	if err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	context.Status(204)
}
