package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Log struct {
	ID              uuid.UUID `gorm:"primaryKey"`
	ApplicationName string    `gorm:"index"`
	Level           string    `gorm:"index"`
	Message         string
	StatusCode      int
	Metadata        datatypes.JSON
	CreatedAt       time.Time
}

func NewLogFromPayload(logPayload *LogPayload) (*Log, error) {
	jsonBytes, err := json.Marshal(logPayload.Metadata)
	if err != nil {
		return nil, err
	}

	metadataJSON := datatypes.JSON(jsonBytes)
	return &Log{
		ID:              uuid.New(),
		ApplicationName: logPayload.ApplicationName,
		Level:           logPayload.Level,
		StatusCode:      logPayload.StatusCode,
		Message:         logPayload.Message,
		Metadata:        metadataJSON,
	}, nil
}
