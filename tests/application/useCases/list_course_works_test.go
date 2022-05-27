package usecases_test

import (
	"testing"

	"github.com/monitoring-go/src/application/protocols/repositories"
	"github.com/monitoring-go/src/domain/entities"
	"github.com/stretchr/testify/require"
)

type ListCourseWorksRepositoryStub struct {
	CallsCount int
}

func (repo *ListCourseWorksRepositoryStub) List() []*entities.CourseWork {
	repo.CallsCount++
	return nil
}

func NewListCourseWorksRepositoryStub() *ListCourseWorksRepositoryStub {
	return &ListCourseWorksRepositoryStub{
		CallsCount: 0,
	}
}

type ListCourseWorksService struct {
	repo repositories.ListCourseWorksRepository
}

func NewListCourseWorksService(repo repositories.ListCourseWorksRepository) *ListCourseWorksService {
	return &ListCourseWorksService{
		repo: repo,
	}
}

func (service *ListCourseWorksService) List() []*entities.CourseWork {
	service.repo.List()
	return nil
}

func TestListCourseWorks_ShouldCallRepository(t *testing.T) {
	repo := NewListCourseWorksRepositoryStub()
	listCourseWorksService := NewListCourseWorksService(repo)

	listCourseWorksService.List()

	require.Equal(t, 1, repo.CallsCount)
}
