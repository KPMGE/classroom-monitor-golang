package entities

type Student struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewStudent(name string, email string) *Student {
	return &Student{
		Name:  name,
		Email: email,
	}
}
