package main

import (
	"fibertest/controllers"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	runServer()
}

func runServer() {
	app := fiber.New()
	controllers.AddControllers(app)
	app.Listen(":8000")
}
