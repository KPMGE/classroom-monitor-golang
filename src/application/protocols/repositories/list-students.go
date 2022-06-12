package repositories

import "github.com/monitoring-go/src/domain/entities"

type ListStudentsRepository interface {
	List() ([]*entities.Student, error)
}
