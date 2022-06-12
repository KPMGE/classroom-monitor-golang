package repositories

import "github.com/monitoring-go/src/domain/entities"

type ListStudentsRepository interface {
	ListStudents(courseId string) ([]*entities.Student, error)
}
