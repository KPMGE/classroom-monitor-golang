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
	students, err := controller.service.List()
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
