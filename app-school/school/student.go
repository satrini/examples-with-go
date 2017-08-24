package school

import "fmt"

type student struct {
	Name   string
	Age    int
	Note1  float64
	Note2  float64
	Media  float64
	Status bool
}

func SetStudent(name string, age int, note1 float64, note2 float64) *student {
	s := &student{Name: name, Age: age, Note1: note1, Note2: note2}
	return s
}

func (s *student) SetMedia() {
	s.Media = (s.Note1 + s.Note2) / 2
}

func (s *student) SetStatus() {
	if s.Media < 6 {
		s.Status = false
	} else {
		s.Status = true
	}
}

func (s *student) Show() {
	fmt.Println("Name:", s.Name)
	fmt.Println("note 1:", s.Note1)
	fmt.Println("note 2:", s.Note2)
	fmt.Println("media:", s.Media)

	if s.Status {
		fmt.Println("status: approved")
	} else {
		fmt.Println("status: reproved")
	}
}
