package router

import (
	"github.com/Sahilb315/trello_clone/controller"
	// "github.com/Sahilb315/trello_clone/logic"
	"github.com/Sahilb315/trello_clone/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	api.Get("/", controller.TestAPI)
	api.Post("/login", controller.Login)
	api.Post("/sign_up", controller.SignUp)

	api.Use(middlewares.VerifyToken)
	api.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, from Trello API!")
	})
}
