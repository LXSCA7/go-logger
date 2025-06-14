package models

type LogPayload struct {
	ApplicationName string                 `json:"application_name" validate:"required"`
	Level           string                 `json:"level" validate:"required"`
	StatusCode      int                    `json:"status_code"`
	Message         string                 `json:"message" validate:"required"`
	Metadata        map[string]interface{} `json:"metadata"`
}
