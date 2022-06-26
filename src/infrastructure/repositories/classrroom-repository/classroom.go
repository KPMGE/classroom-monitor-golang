package classrroomrepository

import (
	"github.com/monitoring-go/src/domain/entities"
	"google.golang.org/api/classroom/v1"
)

type ClassroomRepository struct{}

func GetAllStudentSubmissions(srv *classroom.Service, courseId string, courseWorkId string) ([]*entities.Submission, error) {
	r, err := srv.Courses.CourseWork.StudentSubmissions.
		List(courseId, courseWorkId).
		States("TURNED_IN").
		Fields("studentSubmissions.id,studentSubmissions.late,studentSubmissions.userId").
		Do()

	if err != nil {
		return nil, err
	}

	submissions := []*entities.Submission{}

	allStudents, err := GetAllStudents(srv, courseId, []*entities.Student{}, "")
	if err != nil {
		return nil, err
	}

	for _, submission := range r.StudentSubmissions {
		for _, student := range allStudents {
			if student.ID == submission.UserId {
				newSubmission := entities.NewSubmission(submission.Id, submission.Late, student)
				submissions = append(submissions, newSubmission)
			}
		}
	}

	return submissions, nil
}

func GetAllCourseWorks(srv *classroom.Service, courseId string) ([]*entities.CourseWork, error) {
	r, err := srv.Courses.CourseWork.
		List(courseId).
		OrderBy("dueDate asc").
		Fields("courseWork.id,courseWork.title").
		Do()

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

func GetAllStudents(srv *classroom.Service, courseId string, students []*entities.Student, nextPageToken string) ([]*entities.Student, error) {
	response, err := srv.Courses.Students.List(courseId).PageToken(nextPageToken).Do()
	if err != nil {
		return nil, err
	}

	for _, s := range response.Students {
		student := entities.NewStudent(s.Profile.Id, s.Profile.Name.FullName, s.Profile.EmailAddress)
		students = append(students, student)
	}

	if response.NextPageToken != "" {
		return GetAllStudents(srv, courseId, students, response.NextPageToken)
	}

	return students, nil
}

func (repo *ClassroomRepository) ListStudents(courseId string) ([]*entities.Student, error) {
	srv := GetClassroomService()
	students, err := GetAllStudents(srv, courseId, []*entities.Student{}, "")
	return students, err
}

func (repo *ClassroomRepository) ListCourses() ([]*entities.Course, error) {
	srv := GetClassroomService()
	courses, err := GetAllCourses(srv)
	return courses, err
}

func (repo *ClassroomRepository) ListCourseWorks(courseId string) ([]*entities.CourseWork, error) {
	srv := GetClassroomService()
	courseWorks, err := GetAllCourseWorks(srv, courseId)
	return courseWorks, err
}

func NewClassroomRepository() *ClassroomRepository {
	return &ClassroomRepository{}
}
