package main

import (
	"github.com/LXSCA7/go-logger/config"
	"github.com/LXSCA7/go-logger/handlers"
	"github.com/LXSCA7/go-logger/repositories"
	"github.com/LXSCA7/go-logger/routes"
	"github.com/LXSCA7/go-logger/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	vars, err := config.LoadEnvVars()
	config.Validate(vars)
	allowedApps := config.LoadApps(vars.SkipAppValidations)
	db, err := config.ConnectDB(vars)
	if err != nil {
		panic(err.Error())
	}

	repo := repositories.NewGormLoggerRepository(db)
	svc := services.NewLoggerService(repo)
	deps := routes.RouteDependencies{
		App:                app,
		Handler:            handlers.NewLoggerHandler(svc),
		ApiKey:             vars.ApiKey,
		AllowedApps:        allowedApps,
		SkipAppValidations: vars.SkipAppValidations,
	}

	routes.SetupRoutes(deps)
	app.Listen(":" + vars.ApiPort)
}
