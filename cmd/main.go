package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leonardobaranelli/rest-api-backend-app/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server on")
	})

	app.Listen(":3000")
}
