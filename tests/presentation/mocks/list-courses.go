package mocks_test

import "github.com/monitoring-go/src/domain/entities"

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
