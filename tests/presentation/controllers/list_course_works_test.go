package controllers_test

import (
	"errors"
	"testing"

	"github.com/monitoring-go/src/presentation/controllers"
	httpprotocols "github.com/monitoring-go/src/presentation/http-protocols"
	mocks_test "github.com/monitoring-go/tests/presentation/mocks"
	"github.com/stretchr/testify/require"
)

func MakeListCourseWorksControllerSut() (*mocks_test.ListCourseWorksServiceSpy, *controllers.ListCourseWorksController) {
	service := mocks_test.NewListCourseWorksServiceSpy()
	controller := controllers.NewListCourseWorksController(service)
	return service, controller
}

func TestController_ShouldReturnServerErrorIfServiceReturnsError(t *testing.T) {
	service, controller := MakeListCourseWorksControllerSut()
	service.Error = errors.New("service error")

	request := &httpprotocols.HttpRequest{
		Params: "any id",
		Body:   nil,
	}
	httpResponse := controller.Handle(request)

	require.Equal(t, 500, httpResponse.StatusCode)
	require.Equal(t, service.Error, httpResponse.Body)
}

func TestController_ShouldReturnRightDataOnSuccess(t *testing.T) {
	service, controller := MakeListCourseWorksControllerSut()

	request := &httpprotocols.HttpRequest{
		Params: "any id",
		Body:   nil,
	}
	httpResponse := controller.Handle(request)

	require.Equal(t, 200, httpResponse.StatusCode)
	require.Equal(t, service.Output, httpResponse.Body)
}
