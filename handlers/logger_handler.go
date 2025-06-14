package handlers

import (
	"github.com/LXSCA7/go-logger/models"
	"github.com/LXSCA7/go-logger/services"
	"github.com/gofiber/fiber/v2"
)

type LoggerHandler struct {
	loggerService services.LoggerService
}

func NewLoggerHandler(svc services.LoggerService) *LoggerHandler {
	return &LoggerHandler{loggerService: svc}
}

func (h *LoggerHandler) Log(c *fiber.Ctx) error {
	var logPayload models.LogPayload
	if err := c.BodyParser(&logPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}

	err := h.loggerService.CreateLog(&logPayload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusCreated)
}
