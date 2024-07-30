package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nelsonstos/isgchallenge/models"
	"github.com/nelsonstos/isgchallenge/repositories"
	"github.com/nelsonstos/isgchallenge/services"
)

func Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := repositories.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create user"})
	}

	token, err := services.GenerateJWT(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create JWT"})
	}

	return c.JSON(fiber.Map{"token": token})
}
