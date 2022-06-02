package domain_test

import "github.com/monitoring-go/src/domain/entities"

func MakeFakeSubmission() *entities.Submission {
	return &entities.Submission{
		ID:        "any_id",
		StudentId: "any_student_id",
		Late:      false,
		Student:   MakeFakeStudent(),
	}
}
