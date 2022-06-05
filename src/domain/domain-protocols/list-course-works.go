package domainprotocols

import "github.com/monitoring-go/src/domain/entities"

type ListCourseWorksUseCase interface {
	List(courseId string) ([]*entities.CourseWork, error)
}
