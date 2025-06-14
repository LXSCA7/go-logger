package main

import (
	"github.com/LXSCA7/go-logger/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	vars, err := config.LoadEnvVars()
	if err != nil {
		panic(err.Error())
	}

	app.Listen(":" + vars.ApiPort)
}
