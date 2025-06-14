package config

import (
	"errors"
	"log"
	"os"

	"github.com/LXSCA7/go-logger/models"
	"github.com/joho/godotenv"
)

func LoadEnvVars() (*models.EnvVars, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Aviso: arquivo .env não encontrado, usando variáveis de ambiente do sistema.")
		}
	}

	envVars := &models.EnvVars{
		DbHost: os.Getenv("DB_HOST"),
		DbPass: os.Getenv("DB_PASS"),
		DbUser: os.Getenv("DB_USER"),
	}

	if envVars.DbHost == "" {
		return &models.EnvVars{}, errors.New("Required environmennt variable 'DB_HOST' not found.")
	}

	if envVars.DbPass == "" {
		return &models.EnvVars{}, errors.New("Required environmennt variable 'DB_PASS' not found.")
	}

	if envVars.DbUser == "" {
		return &models.EnvVars{}, errors.New("Required environmennt variable 'DB_USER' not found.")
	}

	return envVars, nil
}
