package domainprotocols

import "github.com/monitoring-go/src/domain/entities"

type ListCoursesUseCase interface {
	List() ([]*entities.Course, error)
}
