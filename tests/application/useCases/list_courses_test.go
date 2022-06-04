package usecases_test

import (
	"errors"
	"testing"

	"github.com/monitoring-go/src/application/services"
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

func MakeListCoursesSut() (*ListCoursesRepositoryStub, *services.ListCoursesService) {
	repo := NewListCoursesRepositoryStub()
	sut := services.NewListCoursesService(repo)

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
