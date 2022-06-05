package repositories

import "github.com/monitoring-go/src/domain/entities"

type ListCourseWorksRepository interface {
	List(courseId string) ([]*entities.CourseWork, error)
}
