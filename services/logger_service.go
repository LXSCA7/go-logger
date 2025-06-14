package services

import "github.com/LXSCA7/go-logger/models"

type LoggerService interface {
	CreateLog(logPayload *models.LogPayload) error
	GetLogByAppName(appName string) (*[]models.Log, error)
	ListAllLogs() (*[]models.Log, error)
}
