package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/monitoring-go/src/main/configuration"
)

func main() {
	app := fiber.New()
	configuration.SetupRoutes(app)
	app.Listen(":3333")
}
