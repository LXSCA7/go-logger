package middlewares

import (
	"crypto/subtle"

	"github.com/gofiber/fiber/v2"
)

const apiKeyHeader = "X-API-KEY"

func ApiKeyAuth(apiKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientApiKey := c.Get(apiKeyHeader)

		if subtle.ConstantTimeCompare([]byte(apiKey), []byte(clientApiKey)) != 1 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		return c.Next()
	}
}
