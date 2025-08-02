package views

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func HomeView(c *fiber.Ctx) error {
	return c.SendString("Hallo Saya Rafi")
}

func MarketView(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Market endpoint hit",
	})
}

func NameAgeView(c *fiber.Ctx) error {
	name := c.Params("name")
	age := c.Params("age")
	var message string
	if age == "" {
		message = "Berapa Usianya?"
	} else {
		message = "Mempunyai umur " + age
	}
	return c.JSON(fiber.Map{
		"message": message,
		"name":    name,
		"now":     time.Now().Format("2006-01-02 15:04:05"),
	})
}
