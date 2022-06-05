package controllers

import (
	domainprotocols "github.com/monitoring-go/src/domain/domain-protocols"
	httphelpers "github.com/monitoring-go/src/presentation/http-helpers"
	httpprotocols "github.com/monitoring-go/src/presentation/http-protocols"
)

type ListCourseWorksController struct {
	service domainprotocols.ListCourseWorksUseCase
}

func (controller *ListCourseWorksController) Handle(request *httpprotocols.HttpRequest) *httpprotocols.HttpResponse {
	courseId := request.Params

	if courseId == nil {
		return httphelpers.BadRequest("Course Id not provided!")
	}

	courseWorks, err := controller.service.List(courseId.(string))

	if err != nil {
		return httphelpers.ServerError(err)
	}

	return httphelpers.Ok(courseWorks)
}

func NewListCourseWorksController(service domainprotocols.ListCourseWorksUseCase) *ListCourseWorksController {
	return &ListCourseWorksController{
		service: service,
	}
}
