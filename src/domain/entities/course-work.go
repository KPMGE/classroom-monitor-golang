package entities

type CourseWork struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Submissions []*Submission `json:"submissions"`
}

func NewCourseWork(id string, title string, submissions []*Submission) *CourseWork {
	return &CourseWork{
		ID:          id,
		Title:       title,
		Submissions: submissions,
	}
}
