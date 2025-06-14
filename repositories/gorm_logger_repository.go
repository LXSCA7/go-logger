package repositories

import "github.com/LXSCA7/go-logger/models"

type GormLoggerRepository interface {
	Log(*models.Log) error
	ListAll() (*[]models.Log, error)
	ListByAppName(appName string) (*[]models.Log, error)
}
