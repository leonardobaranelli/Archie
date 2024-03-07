package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/leonardobaranelli/rest-api-backend-app/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Error starting the server: %s", err)
	}
}
