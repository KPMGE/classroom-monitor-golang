package classrroomrepository

import (
	"github.com/monitoring-go/src/domain/entities"
	"google.golang.org/api/classroom/v1"
)

type ClassroomRepository struct{}

func GetStudent(srv *classroom.Service, courseId string, studentId string) (*entities.Student, error) {
	r, err := srv.Courses.Students.Get(courseId, studentId).Do()

	if err != nil {
		return nil, err
	}

	name := r.Profile.Name.FullName
	email := r.Profile.EmailAddress

	student := entities.NewStudent(name, email)

	return student, nil
}

func GetAllStudentSubmissions(srv *classroom.Service, courseId string, courseWorkId string) ([]*entities.Submission, error) {
	r, err := srv.Courses.CourseWork.StudentSubmissions.List(courseId, courseWorkId).Do()

	if err != nil {
		return nil, err
	}

	submissions := []*entities.Submission{}

	for _, s := range r.StudentSubmissions {
		student, err := GetStudent(srv, courseId, s.UserId)

		if err != nil {
			return nil, err
		}

		newSubmission := entities.NewSubmission(s.Id, s.UserId, s.Late, student)
		submissions = append(submissions, newSubmission)
	}

	return submissions, nil
}

func GetAllCourseWorks(srv *classroom.Service, courseId string) ([]*entities.CourseWork, error) {
	r, err := srv.Courses.CourseWork.List(courseId).Do()

	if err != nil {
		return nil, err
	}

	courseWorks := []*entities.CourseWork{}

	for _, c := range r.CourseWork {
		submissions, err := GetAllStudentSubmissions(srv, courseId, c.Id)

		if err != nil {
			return nil, err
		}

		courseWork := entities.NewCourseWork(c.Id, c.Title, submissions)
		courseWorks = append(courseWorks, courseWork)
	}

	return courseWorks, nil
}

func (repo *ClassroomRepository) List(courseId string) ([]*entities.CourseWork, error) {
	srv := GetClassroomService()
	courseWorks, err := GetAllCourseWorks(srv, courseId)
	return courseWorks, err
}

func NewClassroomRepository() *ClassroomRepository {
	return &ClassroomRepository{}
}
