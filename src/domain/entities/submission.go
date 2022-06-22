package entities

type Submission struct {
	ID      string   `json:"id"`
	Late    bool     `json:"late"`
	Student *Student `json:"student"`
}

func NewSubmission(id string, late bool, student *Student) *Submission {
	return &Submission{
		ID:      id,
		Late:    late,
		Student: student,
	}
}
