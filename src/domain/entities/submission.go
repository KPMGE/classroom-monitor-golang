package entities

type Submission struct {
	ID        string
	StudentId string
	Late      bool
	Student   *Student
}

func NewSubmission(id string, studentId string, late bool, student *Student) *Submission {
	return &Submission{
		ID:        id,
		StudentId: studentId,
		Late:      late,
		Student:   student,
	}
}
