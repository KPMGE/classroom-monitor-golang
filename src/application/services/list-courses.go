package services

import (
	"github.com/monitoring-go/src/application/protocols/repositories"
	"github.com/monitoring-go/src/domain/entities"
)

type ListCoursesService struct {
	repo repositories.ListCoursesRepository
}

func (service *ListCoursesService) List() ([]*entities.Course, error) {
	courses, err := service.repo.List()

	if err != nil {
		return nil, err
	}

	return courses, nil
}

func NewListCoursesService(repo repositories.ListCoursesRepository) *ListCoursesService {
	return &ListCoursesService{
		repo: repo,
	}
}
