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

func (h *LoggerHandler) ListAllByAppName(c *fiber.Ctx) error {
	appName := c.Params("appName")
	if appName == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "App not found",
		})
	}

	logs, err := h.loggerService.GetLogByAppName(appName)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"logs": logs,
	})
}

func (h *LoggerHandler) ListApps(c *fiber.Ctx) error {
	apps, err := h.loggerService.ListApps()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"apps": apps,
	})
}
