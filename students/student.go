package students

type Student struct {
	Name   string
	RollNo int
	// subjects []string
}

func (s Student) getName() string {
	return s.Name
}
