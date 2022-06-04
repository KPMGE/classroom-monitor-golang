package controllers_test

import (
	"errors"
	"testing"

	"github.com/monitoring-go/src/domain/entities"
	"github.com/monitoring-go/src/presentation/controllers"
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

func MakeListCoursesSut() (*ListCoursesServiceStub, *controllers.ListCoursesController) {
	service := NewListCoursesServiceStub()
	sut := controllers.NewListCoursesController(service)
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
