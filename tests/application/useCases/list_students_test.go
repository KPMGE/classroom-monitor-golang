package usecases_test

import (
	"testing"

	"github.com/monitoring-go/src/domain/entities"
	"github.com/stretchr/testify/require"
)

type ListStudentsRepository interface {
	List() ([]*entities.Student, error)
}

type ListStudentsRepositoryMock struct {
	CallsCount int
}

func (repo *ListStudentsRepositoryMock) List() ([]*entities.Student, error) {
	repo.CallsCount++
	return nil, nil
}

func NewListStudentsRepositoryMock() *ListStudentsRepositoryMock {
	return &ListStudentsRepositoryMock{
		CallsCount: 0,
	}
}

type ListStudentsService struct {
	repo ListStudentsRepository
}

func (service *ListStudentsService) List() ([]*entities.Student, error) {
	service.repo.List()

	return nil, nil
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
