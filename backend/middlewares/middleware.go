package middlewares

import (
	"fmt"
	"os"

	"github.com/Sahilb315/trello_clone/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.SendResponse(c, fiber.StatusUnauthorized, "Unauthorized", nil)
	}

	tokenString := authHeader[len("Bearer "):]
	secret := []byte(os.Getenv("SECRET"))
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return secret, nil
	})
	if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid token",
        })
    }

    // Check if the token is valid
    if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		fmt.Println("OK", ok, "token valid", token.Valid)
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid token",
        })
    }
    return c.Next()

}
