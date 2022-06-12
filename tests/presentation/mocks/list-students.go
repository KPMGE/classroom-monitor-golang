package mocks_test

import (
	"github.com/monitoring-go/src/domain/entities"
	domain_test "github.com/monitoring-go/tests/domain"
)

type ListStudentsServiceMock struct {
	Output []*entities.Student
	Error  error
}

func (service *ListStudentsServiceMock) List() ([]*entities.Student, error) {
	return service.Output, service.Error
}

func NewListStudentsServiceMock() *ListStudentsServiceMock {
	return &ListStudentsServiceMock{
		Output: []*entities.Student{domain_test.MakeFakeStudent()},
		Error:  nil,
	}
}
