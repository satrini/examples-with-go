package main

import "app/school"

func main() {
	teacher := school.SetTeacher("John Doe", 44, "computer science")
	teacher.Show()

	student := school.SetStudent("Marie Doe", 22, 8.0, 9.5)
	student.SetMedia()
	student.SetStatus()
	student.Show()
}
