package middleware

import (
	"os"

	jwtware "github.com/gofiber/jwt/v3"

	"github.com/gofiber/fiber/v2"
)

func AuthRequired() func(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
		SigningKey: []byte(os.Getenv("SECRET_WEB_KEY")),
	})
}
