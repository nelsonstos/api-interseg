package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nelsonstos/isgchallenge/cmd"
	"github.com/nelsonstos/isgchallenge/config"
	"github.com/nelsonstos/isgchallenge/routes"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	cmd.Execute()

	// Inicializa el validador
	config.InitValidation()

	config.ConnectDB()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
