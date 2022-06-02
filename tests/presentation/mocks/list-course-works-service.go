package mocks_test

import (
	"github.com/monitoring-go/src/domain/entities"
	domain_test "github.com/monitoring-go/tests/domain"
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
