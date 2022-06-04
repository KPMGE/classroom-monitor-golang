package factories

import (
	"github.com/monitoring-go/src/application/services"
	clasroomlistcourses "github.com/monitoring-go/src/infrastructure/repositories/classrroom-repository/clasroom-list-courses"
	"github.com/monitoring-go/src/presentation/controllers"
	presentationprotocols "github.com/monitoring-go/src/presentation/presentation-protocols"
)

func MakeListCoursesController() presentationprotocols.Controller {
	repo := clasroomlistcourses.NewListCoursesClassroomRepository()
	service := services.NewListCoursesService(repo)
	controller := controllers.NewListCoursesController(service)
	return controller
}
