package students

import (
	"strings"
)

type Student struct {
	Name   string
	RollNo int
	// subjects []string
}

func (s Student) GetName() string {
	return s.Name
}

func (s Student) GetRoll() int {
	return s.RollNo
}

func (s Student) GetFirstName() string {
	names := strings.Split(s.Name, " ")
	return names[0]
}

func (s Student) GetLastname() string {
	names := strings.Split(s.Name, " ")
	if len(names) > 2 {
		name_str := ""
		for _, name := range names[1:] {
			name_str += name
		}
		return name_str
	}
	return names[len(names)-1]
}
