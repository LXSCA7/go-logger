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
	if err != nil {
		panic(err.Error())
	}

	db, err := config.ConnectDB(vars)
	if err != nil {
		panic(err.Error())
	}

	repo := repositories.NewGormLoggerRepository(db)
	svc := services.NewLoggerService(repo)
	deps := routes.RouteDependencies{
		App:     app,
		Handler: handlers.NewLoggerHandler(svc),
		ApiKey:  vars.ApiKey,
	}

	routes.SetupRoutes(deps)
	app.Listen(":" + vars.ApiPort)
}
