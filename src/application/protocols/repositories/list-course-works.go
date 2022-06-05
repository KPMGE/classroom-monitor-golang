package repositories

import "github.com/monitoring-go/src/domain/entities"

type ListCourseWorksRepository interface {
	ListCourseWorks(courseId string) ([]*entities.CourseWork, error)
}
