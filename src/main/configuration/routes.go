package configuration

// TODO: create adapters to separate the fiber context

import (
	"github.com/gofiber/fiber/v2"
	"github.com/monitoring-go/src/main/factories"
)

func SetupRoutes(app *fiber.App) *fiber.Router {
	api := app.Group("/api")

	api.Get("/course-works", func(c *fiber.Ctx) error {
		controller := factories.MakeListCourseWorksController()
		httpResponse := controller.Handle(nil)
		return c.Status(httpResponse.StatusCode).JSON(httpResponse.Body)
	})

	api.Get("/courses", func(c *fiber.Ctx) error {
		controller := factories.MakeListCoursesController()
		httpResponse := controller.Handle(nil)
		return c.Status(httpResponse.StatusCode).JSON(httpResponse.Body)
	})

	return &api
}
