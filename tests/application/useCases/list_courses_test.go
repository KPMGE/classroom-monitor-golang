package usecases_test

import (
	"errors"
	"testing"

	"github.com/monitoring-go/src/application/protocols/repositories"
	"github.com/monitoring-go/src/domain/entities"
	"github.com/stretchr/testify/require"
)

type ListCoursesRepositoryStub struct {
	Output []*entities.Course
	Error  error
}

func (repo *ListCoursesRepositoryStub) List() ([]*entities.Course, error) {
	return repo.Output, repo.Error
}

func NewListCoursesRepositoryStub() *ListCoursesRepositoryStub {
	return &ListCoursesRepositoryStub{
		Output: []*entities.Course{},
		Error:  nil,
	}
}

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

func MakeListCoursesSut() (*ListCoursesRepositoryStub, *ListCoursesService) {
	repo := NewListCoursesRepositoryStub()
	sut := NewListCoursesService(repo)

	return repo, sut
}

func TestListCourses_ShouldReturnErrorIfRepositoryRetunsError(t *testing.T) {
	repo, sut := MakeListCoursesSut()
	repo.Error = errors.New("repo error")

	courses, err := sut.List()

	require.Nil(t, courses)
	require.Equal(t, repo.Error, err)
}

func TestListCourses_ShouldReturnValidListOfCourses(t *testing.T) {
	repo, sut := MakeListCoursesSut()

	courses, err := sut.List()

	require.Nil(t, err)
	require.Equal(t, repo.Output, courses)
}
