package models

import (
	"time"

	"gorm.io/datatypes"
)

type Log struct {
	ID              uint   `gorm:"primaryKey"`
	ApplicationName string `gorm:"index"`
	Level           string `gorm:"index"`
	Message         string
	StatusCode      int
	Metadata        datatypes.JSON
	CreatedAt       time.Time
}
