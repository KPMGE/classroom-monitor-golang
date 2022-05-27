package entities

type Student struct {
	Name  string
	Email string
}

func NewStudent(name string, email string) *Student {
	return &Student{
		Name:  name,
		Email: email,
	}
}
