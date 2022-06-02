package mocks_test

import (
	"github.com/monitoring-go/src/domain/entities"
	domain_test "github.com/monitoring-go/tests/domain"
)

type ListCourseWorksRepositoryStub struct {
	CallsCount int
	Output     []*entities.CourseWork
	Error      error
}

func (repo *ListCourseWorksRepositoryStub) List() ([]*entities.CourseWork, error) {
	repo.CallsCount++
	return repo.Output, repo.Error
}

func NewListCourseWorksRepositoryStub() *ListCourseWorksRepositoryStub {
	return &ListCourseWorksRepositoryStub{
		Output:     []*entities.CourseWork{domain_test.MakeFakeCourseWork()},
		Error:      nil,
		CallsCount: 0,
	}
}
