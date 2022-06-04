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
	courses, err := service.repo.List()

	if err != nil {
		return nil, err
	}

	return courses, nil
}

func NewListCoursesService(repo ListCoursesRepository) *ListCoursesService {
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
