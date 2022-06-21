package entities

type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewStudent(id string, name string, email string) *Student {
	return &Student{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
