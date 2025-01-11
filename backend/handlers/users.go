package handlers

import (
	"fmt"
	"log"

	"github.com/dev-ruchi/user-management/backend/app"
	"github.com/dev-ruchi/user-management/backend/models"

	"github.com/gofiber/fiber/v2"
	
)


func HandleAddUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	query := `
        INSERT INTO users (name, email, date_of_birth)
        VALUES ($1, $2, $3)
        RETURNING id, name, email, date_of_birth`

	err := app.Db.QueryRow(query, user.Name, user.Email, user.DateOfBirth).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.DateOfBirth,
	)

	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}

	return c.Status(201).JSON(user)

}

func HandleFetchUsers(c *fiber.Ctx) error {
	rows, err := app.Db.Query("SELECT * FROM users")

	if err != nil {
        fmt.Println(err)
        return c.Status(500).JSON(fiber.Map{
            "message": "Something went wrong",
        })
    }

	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.DateOfBirth); err != nil {

			log.Fatal(err)

			return c.Status(500).JSON(fiber.Map{
				"message": "Something went wrong",
			})
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {

		log.Fatal(err)

		return c.Status(500).JSON(fiber.Map{
            "message": "Something went wrong",
        })
	}

	 return c.Status(200).JSON(users)

}

func HandleUpdateUsers(c *fiber.Ctx) error {
	// Get the user ID from URL parameters
	id := c.Params("id")

	var user models.User

	if err := c.BodyParser(&user); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "message": "Bad request",
        })
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
        return c.Status(500).JSON(fiber.Map{
            "message": "Something went wrong",
        })
    }

	// Respond with the updated user details
	return c.Status(200).JSON(user)
}

func HandleDeleteUsers(c *fiber.Ctx) error {

	query := `
      DELETE FROM users WHERE id=$1;`

	  result, err := app.Db.Exec(query, c.Params("id"))
	  if err != nil {
		  fmt.Println(err)
		  return c.Status(500).JSON(fiber.Map{
			  "message": "Something went wrong",
		  })
	  }
  
	  // Check if any rows were affected
	  rowsAffected, err := result.RowsAffected()
	  if err != nil {
		  fmt.Println(err)
		  return c.Status(500).JSON(fiber.Map{
			  "message": "Failed to get affected rows",
		  })
	  }
  
	  if rowsAffected == 0 {
		  return c.Status(404).JSON(fiber.Map{
			  "message": "User not found",
		  })
	  }
  

	return c.Status(204).Send(nil)
}
