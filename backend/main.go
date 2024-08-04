package main

import (
	"log"
	"os"

	"github.com/Sahilb315/trello_clone/database"
	router "github.com/Sahilb315/trello_clone/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	dbErr := database.PostgressConnection(config)

	if err != nil {
		log.Fatal("Error connecting to database", dbErr)
	}

	app := fiber.New()
	app.Use(CORSHandler)
	router.SetupRoutes(app)

	error := app.Listen(":" + os.Getenv("PORT"))

	if error != nil {
		log.Fatal("Error starting server", error)
	}
}

func CORSHandler(c *fiber.Ctx) error {
	allowedOrigins := []string{"http://localhost:3000"}
	origin := c.Get("Origin")

	// Check if the origin is in the allowed list
	for _, o := range allowedOrigins {
		if o == origin {
			c.Set("Access-Control-Allow-Origin", origin)
			break
		}
	}

	c.Set("Access-Control-Allow-Methods", "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS")
	c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
	c.Set("Access-Control-Allow-Credentials", "true")

	// Handle preflight request
	if c.Method() == "OPTIONS" {
		return c.SendStatus(fiber.StatusNoContent)
	}

	return c.Next()
}