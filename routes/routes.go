package routes

import (
	"github.com/LXSCA7/go-logger/handlers"
	"github.com/LXSCA7/go-logger/middlewares"
	"github.com/gofiber/fiber/v2"
)

type RouteDependencies struct {
	App         *fiber.App
	Handler     *handlers.LoggerHandler
	ApiKey      string
	AllowedApps []string
}

func SetupRoutes(deps RouteDependencies) {
	deps.App.Get("/", middlewares.ApiKeyAuth(deps.ApiKey), func(c *fiber.Ctx) error {
		return c.JSON("hello, world!")
	})

	deps.App.Get("/logs", middlewares.ApiKeyAuth(deps.ApiKey))
	deps.App.Get("/logs/:appName", middlewares.ApiKeyAuth(deps.ApiKey), deps.Handler.ListAllByAppName)

	deps.App.Post("/log", middlewares.ApiKeyAuth(deps.ApiKey), middlewares.LoggerAuth(deps.AllowedApps), deps.Handler.Log)
}
