package controllers

import (
	"database/sql"
	"log"
	"main/database"
	"main/models"

	"github.com/gofiber/fiber/v2"
)
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "wrong format"})
	}

	_, err := database.DB.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Username, user.Email)
	if err != nil {
		log.Println("error to add user:", err)
		return c.Status(500).JSON(fiber.Map{"error": "there is an error to create new user"})
	}
	
	return c.Status(201).JSON(fiber.Map{"message": "new user created successfully"})
}

func GetUsers(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "error to get users"})
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "error to read users data"})
		}
		users = append(users, user)
	}
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	err := database.DB.QueryRow("SELECT id, name, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Username, &user.Email)
	if err == sql.ErrNoRows {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
		} else if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "error to receive users data"})
		}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "data format is wrong"})
	}

	_, err := database.DB.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", user.Username, user.Email, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "error to update user"})
	}

	return c.JSON(fiber.Map{"message": "user updated successfully"})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := database.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "error to delete user"})
	}

	return c.JSON(fiber.Map{"message": "user deleted successfully"})
}