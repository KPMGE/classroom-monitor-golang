package usecases_test

import (
	"testing"

	"github.com/monitoring-go/src/domain/entities"
	domain_test "github.com/monitoring-go/tests/domain"
	"github.com/stretchr/testify/require"
)

type ListStudentsRepository interface {
	List() ([]*entities.Student, error)
}

type ListStudentsRepositoryMock struct {
	Output     []*entities.Student
	CallsCount int
}

func (repo *ListStudentsRepositoryMock) List() ([]*entities.Student, error) {
	repo.CallsCount++
	return repo.Output, nil
}

func NewListStudentsRepositoryMock() *ListStudentsRepositoryMock {
	return &ListStudentsRepositoryMock{
		CallsCount: 0,
		Output:     []*entities.Student{domain_test.MakeFakeStudent()},
	}
}

type ListStudentsService struct {
	repo ListStudentsRepository
}

func (service *ListStudentsService) List() ([]*entities.Student, error) {
	students, _ := service.repo.List()
	return students, nil
}

func NewListStudentService(repo ListStudentsRepository) *ListStudentsService {
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
