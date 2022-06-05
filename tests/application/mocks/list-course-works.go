package mocks_test

import (
	"github.com/monitoring-go/src/domain/entities"
	domain_test "github.com/monitoring-go/tests/domain"
)

type ListCourseWorksRepositorySpy struct {
	Input      string
	CallsCount int
	Output     []*entities.CourseWork
	Error      error
}

func (repo *ListCourseWorksRepositorySpy) List(courseId string) ([]*entities.CourseWork, error) {
	repo.Input = courseId
	repo.CallsCount++
	return repo.Output, repo.Error
}

func NewListCourseWorksRepositorySpy() *ListCourseWorksRepositorySpy {
	return &ListCourseWorksRepositorySpy{
		Output:     []*entities.CourseWork{domain_test.MakeFakeCourseWork()},
		Error:      nil,
		CallsCount: 0,
	}
}
