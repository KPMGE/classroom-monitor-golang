package repositories

import "github.com/monitoring-go/src/domain/entities"

type ListCoursesRepository interface {
	List() ([]*entities.Course, error)
}
