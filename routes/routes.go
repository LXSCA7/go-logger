package routes

import (
	"github.com/LXSCA7/go-logger/handlers"
	"github.com/gofiber/fiber/v2"
)

type RouteDependencies struct {
	App     *fiber.App
	Handler *handlers.LoggerHandler
}

func SetupRoutes(deps RouteDependencies) {
	deps.App.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("hello, world!")
	})

	deps.App.Post("/log", deps.Handler.Log)
}
