package repositories

import "github.com/monitoring-go/src/domain/entities"

type ListCoursesRepository interface {
	ListCourses() ([]*entities.Course, error)
}
