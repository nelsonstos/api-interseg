package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/nelsonstos/isgchallenge/controllers"
	"github.com/nelsonstos/isgchallenge/middleware"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api/v1")

	//Matrix
	api.Post("matrix", controllers.Create)

	//Statistics
	api.Post("/matrix/stats", controllers.CreateStatistic)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("¡Welcome to the API!")
	})

	api.Use(middleware.JWTProtected()) // Proteger rutas siguientes con JWT
}
