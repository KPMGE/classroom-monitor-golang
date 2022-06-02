package domainprotocols

import "github.com/monitoring-go/src/domain/entities"

type ListCourseWorksUseCase interface {
	List() ([]*entities.CourseWork, error)
}
