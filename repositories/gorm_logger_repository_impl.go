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
