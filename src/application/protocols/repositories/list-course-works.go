package repositories

import "github.com/monitoring-go/src/domain/entities"

type ListCourseWorksRepository interface {
	List() []*entities.CourseWork
}
