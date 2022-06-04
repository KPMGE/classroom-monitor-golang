package usecases_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type Course struct {
	ID    string
	Title string
}

type ListCoursesRepository interface {
	List() ([]*Course, error)
}

type ListCoursesRepositoryStub struct {
	Output []*Course
	Error  error
}

func (repo *ListCoursesRepositoryStub) List() ([]*Course, error) {
	return repo.Output, repo.Error
}

func NewListCoursesRepositoryStub() *ListCoursesRepositoryStub {
	return &ListCoursesRepositoryStub{
		Output: []*Course{},
		Error:  nil,
	}
}

type ListCoursesUseCase interface {
	List() (*Course, error)
}

type ListCoursesService struct {
	repo ListCoursesRepository
}

func (service *ListCoursesService) List() ([]*Course, error) {
	_, err := service.repo.List()
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func NewListCoursesService(repo ListCoursesRepository) *ListCoursesService {
	return &ListCoursesService{
		repo: repo,
	}
}

func TestListCourses_ShouldReturnErrorIfRepositoryRetunsError(t *testing.T) {
	repo := NewListCoursesRepositoryStub()
	repo.Error = errors.New("repo error")

	sut := NewListCoursesService(repo)

	courses, err := sut.List()

	require.Nil(t, courses)
	require.Equal(t, repo.Error, err)
}
