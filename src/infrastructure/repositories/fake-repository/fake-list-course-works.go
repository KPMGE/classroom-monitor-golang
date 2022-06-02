package fakerepository

import (
	"github.com/monitoring-go/src/domain/entities"
	domain_test "github.com/monitoring-go/tests/domain"
)

type FakeListCourseWorksRepository struct{}

func (fakeRepo *FakeListCourseWorksRepository) List() ([]*entities.CourseWork, error) {
	courseWorks := []*entities.CourseWork{domain_test.MakeFakeCourseWork(), domain_test.MakeFakeCourseWork()}
	return courseWorks, nil
}

func NewFakeListCourseWorksRepository() *FakeListCourseWorksRepository {
	return &FakeListCourseWorksRepository{}
}
