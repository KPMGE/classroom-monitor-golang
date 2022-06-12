package usecases_test

import (
	"errors"
	"testing"

	"github.com/monitoring-go/src/application/services"
	mocks_test "github.com/monitoring-go/tests/application/mocks"
	"github.com/stretchr/testify/require"
)

func MakeListStudentsSut() (*mocks_test.ListStudentsRepositoryMock, *services.ListStudentsService) {
	repo := mocks_test.NewListStudentsRepositoryMock()
	sut := services.NewListStudentService(repo)
	return repo, sut
}

func TestListStudents_ShouldCallRepositoyOnlyOnce(t *testing.T) {
	repo, sut := MakeListStudentsSut()
	sut.List()

	require.Equal(t, 1, repo.CallsCount)
}

func TestListStudents_ShouldReturnAValidStudentList(t *testing.T) {
	repo, sut := MakeListStudentsSut()

	students, err := sut.List()

	require.Nil(t, err)
	require.Equal(t, repo.Output, students)
}

func TestListStudents_ShouldReturnErrorWhenRepositoryReturnsError(t *testing.T) {
	repo, sut := MakeListStudentsSut()
	repo.Error = errors.New("repo error")

	students, err := sut.List()

	require.Nil(t, students)
	require.Equal(t, repo.Error, err)
}
