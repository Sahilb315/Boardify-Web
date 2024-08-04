package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
	StatusCode  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	resp := Response{
		StatusCode:  status,
		Message: message,
		Data:    data,
	}
	return c.Status(status).JSON(resp)
}
