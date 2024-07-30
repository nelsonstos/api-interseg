package controllers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nelsonstos/isgchallenge/config"
	"github.com/nelsonstos/isgchallenge/dto"
	"github.com/nelsonstos/isgchallenge/services"
	"github.com/nelsonstos/isgchallenge/utils"
)

func Index(c *fiber.Ctx) error {
	fmt.Println("Esto es un test para implementar la matriz de orden 3 por 3")
	return c.JSON(fiber.Map{"message": "Welcome to the ISG Challenge API!"})
}

func Create(c *fiber.Ctx) error {
	var req dto.MatrixRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Check data integrity
	if err := config.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Validar la matriz
	if err := validateMatrix(req.Matrix); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	matrix, _ := services.CalculateQR(req)

	return utils.SuccessResponse(c, http.StatusCreated, matrix, "Factorization and register successfully!")
}

// validateMatrix valida la matriz recibida
func validateMatrix(matrix [][]float64) error {
	if len(matrix) == 0 {
		return fmt.Errorf("matrix is empty")
	}

	rowLen := len(matrix[0])
	for _, row := range matrix {
		if len(row) != rowLen {
			return fmt.Errorf("matrix rows have different lengths")
		}
	}

	if len(matrix) != rowLen {
		return fmt.Errorf("matrix is not square")
	}

	return nil
}
