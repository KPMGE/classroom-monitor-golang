package entities

type Submission struct {
	ID        string   `json:"id"`
	StudentId string   `json:"studentId"`
	Late      bool     `json:"late"`
	Student   *Student `json:"student"`
}

func NewSubmission(id string, studentId string, late bool, student *Student) *Submission {
	return &Submission{
		ID:        id,
		StudentId: studentId,
		Late:      late,
		Student:   student,
	}
}
