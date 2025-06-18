package services

import (
	"github.com/LXSCA7/go-logger/models"
	"github.com/LXSCA7/go-logger/repositories"
)

type loggerServiceImpl struct {
	repo repositories.GormLoggerRepository
}

func NewLoggerService(repo repositories.GormLoggerRepository) LoggerService {
	return &loggerServiceImpl{repo: repo}
}

func (l *loggerServiceImpl) CreateLog(logPayload *models.LogPayload) error {
	log, err := models.NewLogFromPayload(logPayload)
	if err != nil {
		return err
	}

	return l.repo.Log(log)
}

func (l *loggerServiceImpl) GetLogByAppName(appName string) (*[]models.Log, error) {
	return l.repo.ListByAppName(appName)
}

func (l *loggerServiceImpl) ListAllLogs() (*[]models.Log, error) {
	return l.repo.ListAll()
}

func (l *loggerServiceImpl) ListApps() ([]string, error) {
	return l.repo.ListAllApps()
}
