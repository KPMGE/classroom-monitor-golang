package usecases_test

import (
	"errors"
	"testing"

	"github.com/monitoring-go/src/application/protocols/repositories"
	"github.com/monitoring-go/src/domain/entities"
	domain_test "github.com/monitoring-go/tests/domain"
	"github.com/stretchr/testify/require"
)

type ListStudentsRepositoryMock struct {
	Output     []*entities.Student
	Error      error
	CallsCount int
}

func (repo *ListStudentsRepositoryMock) List() ([]*entities.Student, error) {
	repo.CallsCount++
	return repo.Output, repo.Error
}

func NewListStudentsRepositoryMock() *ListStudentsRepositoryMock {
	return &ListStudentsRepositoryMock{
		CallsCount: 0,
		Output:     []*entities.Student{domain_test.MakeFakeStudent()},
		Error:      nil,
	}
}

type ListStudentsService struct {
	repo repositories.ListStudentsRepository
}

func (service *ListStudentsService) List() ([]*entities.Student, error) {
	students, err := service.repo.List()
	if err != nil {
		return nil, err
	}
	return students, nil
}

func NewListStudentService(repo repositories.ListStudentsRepository) *ListStudentsService {
	return &ListStudentsService{
		repo: repo,
	}
}

func MakeListStudentsSut() (*ListStudentsRepositoryMock, *ListStudentsService) {
	repo := NewListStudentsRepositoryMock()
	sut := NewListStudentService(repo)
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
