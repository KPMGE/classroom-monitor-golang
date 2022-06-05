package mocks_test

import "github.com/monitoring-go/src/domain/entities"

type ListCoursesRepositoryStub struct {
	Output []*entities.Course
	Error  error
}

func (repo *ListCoursesRepositoryStub) ListCourses() ([]*entities.Course, error) {
	return repo.Output, repo.Error
}

func NewListCoursesRepositoryStub() *ListCoursesRepositoryStub {
	return &ListCoursesRepositoryStub{
		Output: []*entities.Course{},
		Error:  nil,
	}
}
