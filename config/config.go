package config

import (
	"encoding/json"
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
		ApiPort:            os.Getenv("API_PORT"),
		ApiKey:             os.Getenv("API_KEY"),
		DbHost:             os.Getenv("DB_HOST"),
		DbUser:             os.Getenv("DB_USER"),
		DbPass:             os.Getenv("DB_PASS"),
		DbName:             os.Getenv("DB_NAME"),
		DbPort:             os.Getenv("DB_PORT"),
		DbTimeZone:         os.Getenv("DB_TIMEZONE"),
		SkipAppValidations: os.Getenv("SKIP_APP_VALIDATIONS") == "true",
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
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", vars.DbHost, vars.DbUser, vars.DbPass, vars.DbName, vars.DbPort, vars.DbTimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = runMigrations(db)
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

func runMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&models.Log{})
}

func LoadApps(skip bool) []string {
	if skip {
		return []string{}
	}

	var payload struct {
		Apps []string
	}

	file, err := os.ReadFile("apps.json")
	if err != nil {
		panic(fmt.Sprintf(
			"\x1b[31mFailed to read or parse 'apps.json'. This might be due to a missing or corrupted file.\n"+
				"\x1b[33mDid you copy the 'apps.json.example' file to 'apps.json'? Read the README.md for more information.\n"+
				"\x1b[31mError message: %v\x1b[0m",
			err,
		))
	}

	err = json.Unmarshal(file, &payload)
	if err != nil {
		panic(fmt.Sprintf(
			"\x1b[31mFailed to parse 'apps.json'.\n"+
				"\x1b[33mDid you copy the 'apps.json.example' file to 'apps.json'? Read the README.md for more information.\n"+
				"\x1b[31mError message: %v\x1b[0m",
			err,
		))
	}

	if payload.Apps == nil || len(payload.Apps) == 0 {
		panic(fmt.Sprintf(
			"\x1b[31mThe 'apps.json' file does not contain the expected 'apps' array, or it's empty.\n" +
				"\x1b[33mPlease ensure 'apps.json' has a top-level JSON object with an 'apps' field containing a list of strings.\n" +
				"\x1b[33mFor example: {\"apps\": [\"app1\", \"app2\"]}\x1b[0m",
		))
	}
	return payload.Apps
}

func Validate(vars *models.EnvVars) {
	if os.Getenv("APP_ENV") != "production" && vars.SkipAppValidations {
		fmt.Println("\n\x1b[31m" +
			"[SECURITY WARNING]: You are skipping the validation of authorized applications.\n" +
			"This is UNSAFE and CAN NOT be used in PRODUCTION environments.\x1b[0m")
	}

	if os.Getenv("APP_ENV") == "production" && vars.SkipAppValidations {
		panic("\n\x1b[31m" +
			"Error: You are skipping the validation of authorized applications.\n" +
			"This is UNSAFE and CAN NOT be used in PRODUCTION environments.\n" +
			"For fix this error, change the environment variable 'SKIP_APP_VALIDATIONS' to 'false' and set the allowed apps on 'apps.json'\x1b[0m")
	}

	if vars.ApiKey == "" {
		panic("\n\x1b[31m" +
			"Error: Environment variable 'API_KEY' can not be empty.\n" +
			"\x1b[33mDid you copy the '.env.example' file to '.env'? Read the README.md for more information.\n" +
			"\x1b[31mFor fix this error, fill in the environment variable (.env file).\x1b[0m")
	}
}
