package factories

import (
	"github.com/monitoring-go/src/application/services"
	classrroomrepository "github.com/monitoring-go/src/infrastructure/repositories/classrroom-repository"
	"github.com/monitoring-go/src/main/env"
	"github.com/monitoring-go/src/presentation/controllers"
	presentationprotocols "github.com/monitoring-go/src/presentation/presentation-protocols"
)

func MakeListCourseWorksController() presentationprotocols.Controller {
	env := env.GetEnvObject()
	repo := classrroomrepository.NewClassroomRepository(env.CourseId)
	service := services.NewListCourseWorksService(repo)
	controller := controllers.NewListCourseWorksController(service)
	return controller
}
