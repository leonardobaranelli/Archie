package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/leonardobaranelli/rest-api-backend-app/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Error starting the server: %s", err)
	}
}
