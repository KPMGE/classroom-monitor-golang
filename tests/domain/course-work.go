package domain_test

import "github.com/monitoring-go/src/domain/entities"

func MakeFakeCourseWork() *entities.CourseWork {
	return &entities.CourseWork{
		ID:          "any id",
		Title:       "any title",
		Submissions: []*entities.Submission{MakeFakeSubmission()},
	}
}
