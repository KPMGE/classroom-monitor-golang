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
	_, err := controller.service.List()

	if err != nil {
		return httphelpers.ServerError(err)
	}
	return nil
}

func NewListCoursesController(service domainprotocols.ListCoursesUseCase) *ListCoursesController {
	return &ListCoursesController{
		service: service,
	}
}

func TestListCoursesController_ShouldReturnServerErrorIfServiceReturnsError(t *testing.T) {
	service := NewListCoursesServiceStub()
	service.Error = errors.New("service error")

	sut := NewListCoursesController(service)

	httpResponse := sut.Handle(nil)

	require.Equal(t, 500, httpResponse.StatusCode)
	require.Equal(t, service.Error, httpResponse.Body)
}
