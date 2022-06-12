package domainprotocols

import "github.com/monitoring-go/src/domain/entities"

type ListStudentsUseCase interface {
	List() ([]*entities.Student, error)
}
