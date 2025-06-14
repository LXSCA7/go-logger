package config

import (
	"fmt"
	"log"
	"os"

	"github.com/LXSCA7/go-logger/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnvVars() (*models.EnvVars, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Aviso: arquivo .env não encontrado, usando variáveis de ambiente do sistema.")
		}
	}

	envVars := &models.EnvVars{
		ApiPort: os.Getenv("API_PORT"),
		ApiKey:  os.Getenv("API_KEY"),
		DbHost:  os.Getenv("DB_HOST"),
		DbUser:  os.Getenv("DB_USER"),
		DbPass:  os.Getenv("DB_PASS"),
		DbName:  os.Getenv("DB_NAME"),
	}

	if err := validateEnv(envVars.ApiPort, "API_PORT"); err != nil {
		return nil, err
	}
	if err := validateEnv(envVars.ApiKey, "API_KEY"); err != nil {
		return nil, err
	}
	if err := validateEnv(envVars.DbHost, "DB_HOST"); err != nil {
		return nil, err
	}
	if err := validateEnv(envVars.DbUser, "DB_USER"); err != nil {
		return nil, err
	}
	if err := validateEnv(envVars.DbPass, "DB_PASS"); err != nil {
		return nil, err
	}
	if err := validateEnv(envVars.DbName, "DB_NAME"); err != nil {
		return nil, err
	}

	return envVars, nil
}

func ConnectDB(vars *models.EnvVars) (*gorm.DB, error) {
	dsn := fmt.Sprint("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v", vars.DbHost, vars.DbUser, vars.DbPass, vars.DbName, vars.DbPort, vars.DbTimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func validateEnv(value string, key string) error {
	if value == "" {
		return fmt.Errorf("Required environment variable '%s' was not found.", key)
	}
	return nil
}
