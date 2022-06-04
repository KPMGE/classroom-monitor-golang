package main

import (
	"github.com/gofiber/fiber/v2"
	classrroomrepository "github.com/monitoring-go/src/infrastructure/repositories/classrroom-repository"
	"github.com/monitoring-go/src/main/configuration"
)

func main() {
	app := fiber.New()
	classrroomrepository.GetClassroomService()
	configuration.SetupRoutes(app)
	app.Listen(":3333")
}
