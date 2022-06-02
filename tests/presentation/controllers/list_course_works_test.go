package controllers_test

import (
	"errors"
	"testing"

	"github.com/monitoring-go/src/domain/entities"
	domain_test "github.com/monitoring-go/tests/domain"
	"github.com/stretchr/testify/require"
)

type HttpResponse struct {
	Body       any
	StatusCode int
}

type HttpRequest struct {
	Params any
	Body   any
}

func ServerError(err error) *HttpResponse {
	return &HttpResponse{
		Body:       err,
		StatusCode: 500,
	}
}

type ListCourseWorksServiceSpy struct {
	Output []*entities.CourseWork
	Error  error
}

func (service *ListCourseWorksServiceSpy) List() ([]*entities.CourseWork, error) {
	return service.Output, service.Error
}

type ListCourseWorksUseCase interface {
	List() ([]*entities.CourseWork, error)
}

func NewListCourseWorksServiceSpy() *ListCourseWorksServiceSpy {
	return &ListCourseWorksServiceSpy{
		Output: []*entities.CourseWork{domain_test.MakeFakeCourseWork()},
		Error:  nil,
	}
}

type ListCourseWorksController struct {
	service ListCourseWorksUseCase
}

type Controller interface {
	Handle(request *HttpRequest) *HttpResponse
}

func (controller *ListCourseWorksController) Handle(request *HttpRequest) *HttpResponse {
	_, err := controller.service.List()
	if err != nil {
		return ServerError(err)
	}
	return nil
}

func NewListCourseWorksController(service ListCourseWorksUseCase) *ListCourseWorksController {
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
