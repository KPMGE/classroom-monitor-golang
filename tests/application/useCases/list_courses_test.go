package usecases_test

import (
	"errors"
	"testing"

	"github.com/monitoring-go/src/application/services"
	mocks_test "github.com/monitoring-go/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

func MakeListCoursesSut() (*mocks_test.ListCoursesRepositoryStub, *services.ListCoursesService) {
	repo := mocks_test.NewListCoursesRepositoryStub()
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
