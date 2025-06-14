package main

import (
	"github.com/LXSCA7/go-logger/config"
	"github.com/LXSCA7/go-logger/repositories"
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
	// service
	// create dependencies
	// setup routes

	app.Listen(":" + vars.ApiPort)
}
