package controllers_test

import (
	"errors"
	"testing"

	"github.com/monitoring-go/src/presentation/controllers"
	httpprotocols "github.com/monitoring-go/src/presentation/http-protocols"
	mocks_test "github.com/monitoring-go/tests/presentation/mocks"
	"github.com/stretchr/testify/require"
)

func MakeListStudentsControllerSut() (*mocks_test.ListStudentsServiceMock, *controllers.ListStudentsController) {
	service := mocks_test.NewListStudentsServiceMock()
	sut := controllers.NewListStudentsController(service)
	return service, sut
}

func TestListStudentsController_ShouldReturnServerErrorWhenServiceReturnsError(t *testing.T) {
	service, sut := MakeListStudentsControllerSut()
	service.Error = errors.New("service error")

	request := httpprotocols.NewHttpRequest("any_id", nil)
	httpResponse := sut.Handle(request)

	require.Equal(t, 500, httpResponse.StatusCode)
	require.Equal(t, service.Error.Error(), httpResponse.Body)
}

func TestListStudentsController_ShouldReturnBadRequestIfNoIdIsProvided(t *testing.T) {
	service, sut := MakeListStudentsControllerSut()
	service.Error = errors.New("service error")

	request := httpprotocols.NewHttpRequest(nil, nil)
	httpResponse := sut.Handle(request)

	require.Equal(t, 400, httpResponse.StatusCode)
	require.Equal(t, "Course Id not provided!", httpResponse.Body)
}

func TestListStudentsController_ShouldReturnOkaOnSuccess(t *testing.T) {
	service, sut := MakeListStudentsControllerSut()

	request := httpprotocols.NewHttpRequest("any_id", nil)
	httpResponse := sut.Handle(request)

	require.Equal(t, 200, httpResponse.StatusCode)
	require.Equal(t, service.Output, httpResponse.Body)
}
