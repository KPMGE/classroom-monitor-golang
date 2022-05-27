package entities

type CourseWork struct {
	ID          string
	Title       string
	Submissions []*Submission
}

func NewCourseWork(id string, title string, submissions []*Submission) *CourseWork {
	return &CourseWork{
		ID:          id,
		Title:       title,
		Submissions: submissions,
	}
}
