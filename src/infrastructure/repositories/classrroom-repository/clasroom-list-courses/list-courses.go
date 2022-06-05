package clasroomlistcourses

import (
	"github.com/monitoring-go/src/domain/entities"
	classrroomrepository "github.com/monitoring-go/src/infrastructure/repositories/classrroom-repository"
	"google.golang.org/api/classroom/v1"
)

type ListCoursesClassroomRepository struct{}

func GetAllCourses(srv *classroom.Service) ([]*entities.Course, error) {
	// returns only the courses where the user is the teacher.
	response, err := srv.Courses.List().TeacherId("me").Do()

	if err != nil {
		return nil, err
	}

	courses := []*entities.Course{}

	for _, c := range response.Courses {
		course := entities.NewCourse(c.Id, c.Name)
		courses = append(courses, course)
	}

	return courses, nil
}

func (repo *ListCoursesClassroomRepository) List() ([]*entities.Course, error) {
	srv := classrroomrepository.GetClassroomService()
	courses, err := GetAllCourses(srv)
	return courses, err
}

func NewListCoursesClassroomRepository() *ListCoursesClassroomRepository {
	return &ListCoursesClassroomRepository{}
}
