package controllers

import (
	"fibertest/repository"

	"github.com/gofiber/fiber/v2"
)

func AddControllers(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		var aaa = repository.GetCats()
		return c.JSON(aaa)
	})
}