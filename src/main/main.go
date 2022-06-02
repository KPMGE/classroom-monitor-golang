package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/monitoring-go/src/main/config"
)

func main() {
	app := fiber.New()
	config.SetupRoutes(app)
	app.Listen(":3333")
}
