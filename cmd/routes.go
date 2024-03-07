package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leonardobaranelli/rest-api-backend-app/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)
	app.Get("/fact/:id", handlers.GetFact)
	app.Post("/fact", handlers.CreateFact)
	app.Put("/fact/:id", handlers.UpdateFact)
	app.Delete("/fact/:id", handlers.DeleteFact)
}
