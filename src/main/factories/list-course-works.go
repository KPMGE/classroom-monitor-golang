package factories

import (
	"github.com/monitoring-go/src/application/services"
	classrroomrepository "github.com/monitoring-go/src/infrastructure/repositories/classrroom-repository"
	"github.com/monitoring-go/src/presentation/controllers"
	presentationprotocols "github.com/monitoring-go/src/presentation/presentation-protocols"
)

func MakeListCourseWorksController() presentationprotocols.Controller {
	repo := classrroomrepository.NewClassroomRepository()
	service := services.NewListCourseWorksService(repo)
	controller := controllers.NewListCourseWorksController(service)
	return controller
}
