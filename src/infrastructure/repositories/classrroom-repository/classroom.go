package classrroomrepository

import (
	"log"

	"github.com/monitoring-go/src/domain/entities"
	"google.golang.org/api/classroom/v1"
)

type ClassroomRepository struct{}

func GetStudent(srv *classroom.Service, courseId string, studentId string) *entities.Student {
	r, err := srv.Courses.Students.Get(courseId, studentId).Do()

	if err != nil {
		log.Fatal(err)
	}

	name := r.Profile.Name.FullName
	email := r.Profile.EmailAddress

	student := entities.NewStudent(name, email)

	return student
}

func GetAllStudentSubmissions(srv *classroom.Service, courseId string, courseWorkId string) []*entities.Submission {
	r, err := srv.Courses.CourseWork.StudentSubmissions.List(courseId, courseWorkId).Do()

	if err != nil {
		log.Fatal(err)
	}

	submissions := []*entities.Submission{}

	for _, s := range r.StudentSubmissions {
		student := GetStudent(srv, courseId, s.UserId)
		newSubmission := entities.NewSubmission(s.Id, s.UserId, s.Late, student)
		submissions = append(submissions, newSubmission)
	}

	return submissions
}

func GetAllCourseWorks(srv *classroom.Service, courseId string) []*entities.CourseWork {
	r, err := srv.Courses.CourseWork.List(courseId).Do()

	if err != nil {
		log.Fatal(err)
	}

	courseWorks := []*entities.CourseWork{}

	for _, c := range r.CourseWork {
		submissions := GetAllStudentSubmissions(srv, courseId, c.Id)
		courseWork := entities.NewCourseWork(c.Id, c.Title, submissions)
		courseWorks = append(courseWorks, courseWork)
	}

	return courseWorks
}

func (repo *ClassroomRepository) List(courseId string) ([]*entities.CourseWork, error) {
	srv := GetClassroomService()
	courseWorks := GetAllCourseWorks(srv, courseId)
	return courseWorks, nil
}

func NewClassroomRepository() *ClassroomRepository {
	return &ClassroomRepository{}
}
