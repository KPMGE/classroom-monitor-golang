package usecases_test

import (
	"testing"

	"github.com/monitoring-go/src/application/protocols/repositories"
	"github.com/monitoring-go/src/domain/entities"
	"github.com/stretchr/testify/require"
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
		Output:     []*entities.CourseWork{},
		Error:      nil,
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

func (service *ListCourseWorksService) List() ([]*entities.CourseWork, error) {
	return service.repo.List()
}

func TestListCourseWorks_ShouldCallRepository(t *testing.T) {
	repo := NewListCourseWorksRepositoryStub()
	listCourseWorksService := NewListCourseWorksService(repo)

	listCourseWorksService.List()

	require.Equal(t, 1, repo.CallsCount)
}

func TestListCourseWorks_ShouldReturnAValidList(t *testing.T) {
	repo := NewListCourseWorksRepositoryStub()
	listCourseWorksService := NewListCourseWorksService(repo)

	courseWorks, err := listCourseWorksService.List()

	require.Nil(t, err)
	require.Equal(t, courseWorks, repo.Output)
}
