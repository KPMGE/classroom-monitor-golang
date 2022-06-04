package controllers_test

import (
	"errors"
	"testing"

	"github.com/monitoring-go/src/presentation/controllers"
	mocks_test "github.com/monitoring-go/tests/presentation/mocks"
	"github.com/stretchr/testify/require"
)

func MakeListCoursesSut() (*mocks_test.ListCoursesServiceStub, *controllers.ListCoursesController) {
	service := mocks_test.NewListCoursesServiceStub()
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
