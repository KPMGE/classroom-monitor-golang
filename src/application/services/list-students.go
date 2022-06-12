package services

import (
	"github.com/monitoring-go/src/application/protocols/repositories"
	"github.com/monitoring-go/src/domain/entities"
)

type ListStudentsService struct {
	repo repositories.ListStudentsRepository
}

func (service *ListStudentsService) List() ([]*entities.Student, error) {
	students, err := service.repo.ListStudents()
	if err != nil {
		return nil, err
	}
	return students, nil
}

func NewListStudentService(repo repositories.ListStudentsRepository) *ListStudentsService {
	return &ListStudentsService{
		repo: repo,
	}
}
