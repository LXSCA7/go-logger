package repositories

import (
	"github.com/LXSCA7/go-logger/models"
	"gorm.io/gorm"
)

type gormLoggerRepository struct {
	db *gorm.DB
}

func NewGormLoggerRepository(db *gorm.DB) GormLoggerRepository {
	return &gormLoggerRepository{db: db}
}

func (g *gormLoggerRepository) ListAll() (*[]models.Log, error) {
	panic("unimplemented")
}

func (g *gormLoggerRepository) Log(log *models.Log) error {
	return g.db.Create(log).Error
}

func (g *gormLoggerRepository) ListByAppName(appName string) (*[]models.Log, error) {
	// TO-DO: pagination
	var logs []models.Log
	result := g.db.Where("application_name = ?", appName).Find(&logs)
	if result.Error != nil {
		return nil, result.Error
	}

	return &logs, nil
}

func (g *gormLoggerRepository) ListAllApps() ([]string, error) {
	var apps []string
	err := g.db.Model(&models.Log{}).Distinct("application_name").Pluck("application_name", &apps).Error
	if err != nil {
		return nil, err
	}

	return apps, nil
}
