package configuration

// TODO: create adapters to separate the fiber context

import (
	"github.com/gofiber/fiber/v2"
	"github.com/monitoring-go/src/main/factories"
	httpprotocols "github.com/monitoring-go/src/presentation/http-protocols"
)

func SetupRoutes(app *fiber.App) *fiber.Router {
	api := app.Group("/api")

	api.Get("/course-works/:courseId", func(c *fiber.Ctx) error {
		courseId := c.Params("courseId")

		controller := factories.MakeListCourseWorksController()
		request := httpprotocols.NewHttpRequest(courseId, nil)
		httpResponse := controller.Handle(request)
		return c.Status(httpResponse.StatusCode).JSON(httpResponse.Body)
	})

	api.Get("/courses", func(c *fiber.Ctx) error {
		controller := factories.MakeListCoursesController()
		httpResponse := controller.Handle(nil)
		return c.Status(httpResponse.StatusCode).JSON(httpResponse.Body)
	})

	return &api
}
