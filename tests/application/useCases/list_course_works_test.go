package usecases_test

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type Student struct {
	Name  string
	Email string
}

type Submission struct {
	ID        string
	StudentId string
	Late      bool
	Student   Student
}

type CourseWork struct {
	ID          string
	Title       string
	Submissions []*Submission
}

type ListCourseWorksRepository interface {
	List() []*CourseWork
}

type ListCourseWorksRepositoryStub struct {
	CallsCount int
}

func (repo *ListCourseWorksRepositoryStub) List() []*CourseWork {
	repo.CallsCount++
	return nil
}

func NewListCourseWorksRepositoryStub() *ListCourseWorksRepositoryStub {
	return &ListCourseWorksRepositoryStub{
		CallsCount: 0,
	}
}

type ListCourseWorksService struct {
	repo ListCourseWorksRepository
}

func NewListCourseWorksService(repo ListCourseWorksRepository) *ListCourseWorksService {
	return &ListCourseWorksService{
		repo: repo,
	}
}

func (service *ListCourseWorksService) List() []*CourseWork {
	service.repo.List()
	return nil
}

func TestListCourseWorks_ShouldCallRepository(t *testing.T) {
	repo := NewListCourseWorksRepositoryStub()
	listCourseWorksService := NewListCourseWorksService(repo)

	listCourseWorksService.List()

	require.Equal(t, 1, repo.CallsCount)
}
