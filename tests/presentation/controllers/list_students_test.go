package controllers_test

import (
	"errors"
	"testing"

	"github.com/monitoring-go/src/domain/entities"
	httphelpers "github.com/monitoring-go/src/presentation/http-helpers"
	httpprotocols "github.com/monitoring-go/src/presentation/http-protocols"
	domain_test "github.com/monitoring-go/tests/domain"
	"github.com/stretchr/testify/require"
)

type ListStudentsUseCase interface {
	List() ([]*entities.Student, error)
}

type ListStudentsServiceMock struct {
	Output []*entities.Student
	Error  error
}

func (service *ListStudentsServiceMock) List() ([]*entities.Student, error) {
	return service.Output, service.Error
}

func NewListStudentsServiceMock() *ListStudentsServiceMock {
	return &ListStudentsServiceMock{
		Output: []*entities.Student{domain_test.MakeFakeStudent()},
		Error:  nil,
	}
}

type ListStudentsController struct {
	service ListStudentsUseCase
}

func (controller *ListStudentsController) Handle(request *httpprotocols.HttpRequest) *httpprotocols.HttpResponse {
	_, err := controller.service.List()
	if err != nil {
		return httphelpers.ServerError(err)
	}
	return nil
}

func NewListStudentsController(service ListStudentsUseCase) *ListStudentsController {
	return &ListStudentsController{
		service: service,
	}
}

func TestListStudentsController_ShouldReturnServerErrorWhenServiceReturnsError(t *testing.T) {
	service := NewListStudentsServiceMock()
	sut := NewListStudentsController(service)
	service.Error = errors.New("service error")

	httpResponse := sut.Handle(nil)

	require.Equal(t, 500, httpResponse.StatusCode)
	require.Equal(t, service.Error.Error(), httpResponse.Body)
}
