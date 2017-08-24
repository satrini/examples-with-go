package school

import (
	"fmt"
)

type teacher struct {
	Name   string
	Age    int
	Lesson string
}

// SetTeacher set attributes for struct teacher
func SetTeacher(name string, age int, lesson string) *teacher {
	t := &teacher{Name: name, Age: age, Lesson: lesson}
	return t
}

func (t *teacher) Show() {
	fmt.Println("Name:", t.Name)
	fmt.Println("Age:", t.Age)
	fmt.Println("Lesson:", t.Lesson)
}
