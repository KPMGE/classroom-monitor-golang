package controllers

import (
	domainprotocols "github.com/monitoring-go/src/domain/domain-protocols"
	httphelpers "github.com/monitoring-go/src/presentation/http-helpers"
	httpprotocols "github.com/monitoring-go/src/presentation/http-protocols"
)

type ListCoursesController struct {
	service domainprotocols.ListCoursesUseCase
}

func (controller *ListCoursesController) Handle(request *httpprotocols.HttpRequest) *httpprotocols.HttpResponse {
	courses, err := controller.service.List()

	if err != nil {
		return httphelpers.ServerError(err)
	}

	return httphelpers.Ok(courses)
}

func NewListCoursesController(service domainprotocols.ListCoursesUseCase) *ListCoursesController {
	return &ListCoursesController{
		service: service,
	}
}
