package factories

import (
	"github.com/monitoring-go/src/application/services"
	classrroomrepository "github.com/monitoring-go/src/infrastructure/repositories/classrroom-repository"
	"github.com/monitoring-go/src/presentation/controllers"
	presentationprotocols "github.com/monitoring-go/src/presentation/presentation-protocols"
)

func NewListStudentsController() presentationprotocols.Controller {
	repo := classrroomrepository.NewClassroomRepository()
	service := services.NewListStudentService(repo)
	controller := controllers.NewListStudentsController(service)
	return controller
}
