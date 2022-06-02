package controllers_test

import (
	"errors"
	"testing"

	"github.com/monitoring-go/src/domain/entities"
	"github.com/monitoring-go/src/presentation/controllers"
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

func MakeListCourseWorksControllerSut() (*ListCourseWorksServiceSpy, *controllers.ListCourseWorksController) {
	service := NewListCourseWorksServiceSpy()
	controller := controllers.NewListCourseWorksController(service)
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
