package factories

import (
	"github.com/monitoring-go/src/application/services"
	"github.com/monitoring-go/src/presentation/controllers"
	presentationprotocols "github.com/monitoring-go/src/presentation/presentation-protocols"
	mocks_test "github.com/monitoring-go/tests/application/mocks"
)

func MakeListCoursesController() presentationprotocols.Controller {
	repo := mocks_test.NewListCoursesRepositoryStub()
	service := services.NewListCoursesService(repo)
	controller := controllers.NewListCoursesController(service)
	return controller
}
