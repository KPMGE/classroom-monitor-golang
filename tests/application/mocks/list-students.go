package mocks_test

import (
	"github.com/monitoring-go/src/domain/entities"
	domain_test "github.com/monitoring-go/tests/domain"
)

type ListStudentsRepositoryMock struct {
	Output     []*entities.Student
	Error      error
	CallsCount int
}

func (repo *ListStudentsRepositoryMock) ListStudents(courseId string) ([]*entities.Student, error) {
	repo.CallsCount++
	return repo.Output, repo.Error
}

func NewListStudentsRepositoryMock() *ListStudentsRepositoryMock {
	return &ListStudentsRepositoryMock{
		CallsCount: 0,
		Output:     []*entities.Student{domain_test.MakeFakeStudent()},
		Error:      nil,
	}
}
