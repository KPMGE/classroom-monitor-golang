package configuration

import (
	"github.com/gofiber/fiber/v2"
	"github.com/monitoring-go/src/main/adapters"
	"github.com/monitoring-go/src/main/factories"
)

func SetupRoutes(app *fiber.App) *fiber.Router {
	api := app.Group("/api")

	api.Get("/course-works/:courseId", adapters.FiberRouteAdapter(factories.MakeListCourseWorksController()))
	api.Get("/courses", adapters.FiberRouteAdapter(factories.MakeListCoursesController()))
	api.Get("/students/:courseId", adapters.FiberRouteAdapter(factories.MakeListStudentsController()))

	return &api
}
