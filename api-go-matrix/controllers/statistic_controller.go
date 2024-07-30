package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nelsonstos/isgchallenge/config"
	"github.com/nelsonstos/isgchallenge/dto"
	"github.com/nelsonstos/isgchallenge/services"
	"github.com/nelsonstos/isgchallenge/utils"
)

func CreateStatistic(c *fiber.Ctx) error {
	var req dto.StatisticRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Check data integrity
	if err := config.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	service, _ := services.SaveStatistic(req)
	return utils.SuccessResponse(c, http.StatusCreated, service, "Statistic stored successfully!")
}
