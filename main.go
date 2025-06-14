package main

import (
	"fmt"

	"github.com/LXSCA7/go-logger/config"
)

func main() {
	vars, err := config.LoadEnvVars()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(vars.DbHost, vars.DbPass, vars.DbUser)
}
