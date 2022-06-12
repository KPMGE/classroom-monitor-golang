package controllers

import (
	domainprotocols "github.com/monitoring-go/src/domain/domain-protocols"
	httphelpers "github.com/monitoring-go/src/presentation/http-helpers"
	httpprotocols "github.com/monitoring-go/src/presentation/http-protocols"
)

type ListStudentsController struct {
	service domainprotocols.ListStudentsUseCase
}

func (controller *ListStudentsController) Handle(request *httpprotocols.HttpRequest) *httpprotocols.HttpResponse {
	courseId := request.Params

	if courseId == nil {
		return httphelpers.BadRequest("Course Id not provided!")
	}

	students, err := controller.service.List(courseId.(string))

	if err != nil {
		return httphelpers.ServerError(err)
	}
	return httphelpers.Ok(students)
}

func NewListStudentsController(service domainprotocols.ListStudentsUseCase) *ListStudentsController {
	return &ListStudentsController{
		service: service,
	}
}
