package controllers_test

import (
	"errors"
	"testing"

	domainprotocols "github.com/monitoring-go/src/domain/domain-protocols"
	"github.com/monitoring-go/src/domain/entities"
	httphelpers "github.com/monitoring-go/src/presentation/http-helpers"
	httpprotocols "github.com/monitoring-go/src/presentation/http-protocols"
	"github.com/stretchr/testify/require"
)

type ListCoursesServiceStub struct {
	Output []*entities.Course
	Error  error
}

func NewListCoursesServiceStub() *ListCoursesServiceStub {
	return &ListCoursesServiceStub{
		Output: []*entities.Course{},
		Error:  nil,
	}
}

func (service *ListCoursesServiceStub) List() ([]*entities.Course, error) {
	return service.Output, service.Error
}

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

func MakeListCoursesSut() (*ListCoursesServiceStub, *ListCoursesController) {
	service := NewListCoursesServiceStub()
	sut := NewListCoursesController(service)
	return service, sut
}

func TestListCoursesController_ShouldReturnServerErrorIfServiceReturnsError(t *testing.T) {
	service, sut := MakeListCoursesSut()
	service.Error = errors.New("service error")

	httpResponse := sut.Handle(nil)

	require.Equal(t, 500, httpResponse.StatusCode)
	require.Equal(t, service.Error, httpResponse.Body)
}

func TestListCoursesController_ShouldReturnOkOnSuccess(t *testing.T) {
	service, sut := MakeListCoursesSut()

	httpResponse := sut.Handle(nil)

	require.Equal(t, 200, httpResponse.StatusCode)
	require.Equal(t, service.Output, httpResponse.Body)
}
