package routes

import (
	"github.com/LXSCA7/go-logger/handlers"
	"github.com/LXSCA7/go-logger/middlewares"
	"github.com/gofiber/fiber/v2"
)

type RouteDependencies struct {
	App                *fiber.App
	Handler            *handlers.LoggerHandler
	ApiKey             string
	AllowedApps        []string
	SkipAppValidations bool
}

func SetupRoutes(deps RouteDependencies) {
	deps.App.Use(middlewares.ApiKeyAuth(deps.ApiKey))
	deps.App.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("hello, world!")
	})

	// deps.App.Get("/logs", implement)
	deps.App.Get("/logs/:appName", deps.Handler.ListAllByAppName)

	if deps.SkipAppValidations {
		deps.App.Post("/log", deps.Handler.Log)
	} else {
		deps.App.Post("/log", middlewares.ApplicationsAuth(deps.AllowedApps), deps.Handler.Log)
	}
}
