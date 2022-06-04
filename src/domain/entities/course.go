package entities

type Course struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func NewCourse(id string, title string) *Course {
	return &Course{
		ID:    id,
		Title: title,
	}
}
