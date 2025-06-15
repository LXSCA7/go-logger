package models

type EnvVars struct {
	ApiPort            string
	ApiKey             string
	DbHost             string
	DbUser             string
	DbPass             string
	DbName             string
	DbPort             string
	DbTimeZone         string
	SkipAppValidations bool
}
