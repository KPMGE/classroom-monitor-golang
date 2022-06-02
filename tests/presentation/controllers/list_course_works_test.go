package controllers_test

import (
	"errors"
	"testing"

	domainprotocols "github.com/monitoring-go/src/domain/domain-protocols"
	"github.com/monitoring-go/src/domain/entities"
	httphelpers "github.com/monitoring-go/src/presentation/http-helpers"
	httpprotocols "github.com/monitoring-go/src/presentation/http-protocols"
	domain_test "github.com/monitoring-go/tests/domain"
	"github.com/stretchr/testify/require"
)

type ListCourseWorksServiceSpy struct {
	Output []*entities.CourseWork
	Error  error
}

func (service *ListCourseWorksServiceSpy) List() ([]*entities.CourseWork, error) {
	return service.Output, service.Error
}

func NewListCourseWorksServiceSpy() *ListCourseWorksServiceSpy {
	return &ListCourseWorksServiceSpy{
		Output: []*entities.CourseWork{domain_test.MakeFakeCourseWork()},
		Error:  nil,
	}
}

type ListCourseWorksController struct {
	service domainprotocols.ListCourseWorksUseCase
}

type Controller interface {
	Handle(request *httpprotocols.HttpRequest) *httpprotocols.HttpResponse
}

func (controller *ListCourseWorksController) Handle(request *httpprotocols.HttpRequest) *httpprotocols.HttpResponse {
	courseWorks, err := controller.service.List()

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

func MakeListCourseWorksControllerSut() (*ListCourseWorksServiceSpy, *ListCourseWorksController) {
	service := NewListCourseWorksServiceSpy()
	controller := NewListCourseWorksController(service)
	return service, controller
}

func TestController_ShouldReturnServerErrorIfServiceReturnsError(t *testing.T) {
	service, controller := MakeListCourseWorksControllerSut()
	service.Error = errors.New("service error")

	httpResponse := controller.Handle(nil)

	require.Equal(t, 500, httpResponse.StatusCode)
	require.Equal(t, service.Error, httpResponse.Body)
}

func TestController_ShouldReturnRightDataOnSuccess(t *testing.T) {
	service, controller := MakeListCourseWorksControllerSut()

	httpResponse := controller.Handle(nil)

	require.Equal(t, 200, httpResponse.StatusCode)
	require.Equal(t, service.Output, httpResponse.Body)
}
