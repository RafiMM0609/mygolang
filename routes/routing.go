package routes

import (
	"github.com/gofiber/fiber/v2"
	"market/views"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", views.HomeView)
	app.Post("/market", views.MarketView)
	app.Get("/:name/:age?", views.NameAgeView)
}
