package controller

import (
	"github.com/Sahilb315/trello_clone/database"
	"github.com/gofiber/fiber/v2"
)

func TestAPI(c *fiber.Ctx) error {
	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	db := database.Database.Db
	users := []User{}
	db.Table("users").Find(&users)
	return c.JSON(fiber.Map{
		"message": "API Test Successful",
		"users":   users,
		"count":   len(users),
	})
}
