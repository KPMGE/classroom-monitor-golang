package controllers_test

import (
	"errors"
	"testing"

	"github.com/monitoring-go/src/domain/entities"
	"github.com/monitoring-go/src/presentation/controllers"
	domain_test "github.com/monitoring-go/tests/domain"
	"github.com/stretchr/testify/require"
)

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

func MakeListStudentsControllerSut() (*ListStudentsServiceMock, *controllers.ListStudentsController) {
	service := NewListStudentsServiceMock()
	sut := controllers.NewListStudentsController(service)
	return service, sut
}

func TestListStudentsController_ShouldReturnServerErrorWhenServiceReturnsError(t *testing.T) {
	service, sut := MakeListStudentsControllerSut()
	service.Error = errors.New("service error")

	httpResponse := sut.Handle(nil)

	require.Equal(t, 500, httpResponse.StatusCode)
	require.Equal(t, service.Error.Error(), httpResponse.Body)
}

func TestListStudentsController_ShouldReturnOkaOnSuccess(t *testing.T) {
	service, sut := MakeListStudentsControllerSut()

	httpResponse := sut.Handle(nil)

	require.Equal(t, 200, httpResponse.StatusCode)
	require.Equal(t, service.Output, httpResponse.Body)
}
