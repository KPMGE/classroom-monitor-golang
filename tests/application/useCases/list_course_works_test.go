package usecases_test

import (
	"errors"
	"testing"

	"github.com/monitoring-go/src/application/services"
	mocks_test "github.com/monitoring-go/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

func MakeListCourseWorksSut() (*services.ListCourseWorksService, *mocks_test.ListCourseWorksRepositoryStub) {
	repo := mocks_test.NewListCourseWorksRepositoryStub()
	listCourseWorksService := services.NewListCourseWorksService(repo)
	return listCourseWorksService, repo
}

func TestListCourseWorks_ShouldCallRepository(t *testing.T) {
	listCourseWorksService, repo := MakeListCourseWorksSut()

	listCourseWorksService.List()

	require.Equal(t, 1, repo.CallsCount)
}

func TestListCourseWorks_ShouldReturnAValidList(t *testing.T) {
	listCourseWorksService, repo := MakeListCourseWorksSut()

	courseWorks, err := listCourseWorksService.List()

	require.Nil(t, err)
	require.Equal(t, courseWorks, repo.Output)
}

func TestListCourseWorks_ShouldReturnErrorWhenRepositoryReturnsError(t *testing.T) {
	listCourseWorksService, repo := MakeListCourseWorksSut()
	repo.Error = errors.New("repo error")

	courseWorks, err := listCourseWorksService.List()

	require.Nil(t, courseWorks)
	require.Equal(t, repo.Error, err)
}
