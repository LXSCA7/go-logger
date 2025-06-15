package middlewares

import (
	"crypto/subtle"

	"github.com/LXSCA7/go-logger/models"
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

func LoggerAuth(authorizedApps []string) fiber.Handler {
	var payload models.LogPayload
	return func(c *fiber.Ctx) error {
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "unexpected body",
			})
		}

		for _, appName := range authorizedApps {
			if payload.ApplicationName == appName {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "forbidden access: application not authorized.",
		})
	}
}
