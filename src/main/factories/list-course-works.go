package factories

import (
	"github.com/monitoring-go/src/application/services"
	fakerepository "github.com/monitoring-go/src/infrastructure/repositories/fake-repository"
	"github.com/monitoring-go/src/presentation/controllers"
)

func MakeListCourseWorksController() *controllers.ListCourseWorksController {
	repo := fakerepository.NewFakeListCourseWorksRepository()
	service := services.NewListCourseWorksService(repo)
	controller := controllers.NewListCourseWorksController(service)
	return controller
}
