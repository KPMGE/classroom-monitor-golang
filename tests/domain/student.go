package domain_test

import "github.com/monitoring-go/src/domain/entities"

func MakeFakeStudent() *entities.Student {
	return &entities.Student{
		Name:  "any_name",
		Email: "any_valid_email@gmail.com",
	}
}
