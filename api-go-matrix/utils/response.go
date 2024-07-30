package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Respuesta exitosa
func SuccessResponse(c *fiber.Ctx, status int, data interface{}, message string) error {
	return c.Status(status).JSON(Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// Respuesta de error
func ErrorResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(Response{
		Status:  "error",
		Message: message,
	})
}
