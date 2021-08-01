package main

import (
	"github.com/kirigaikabuto/golanglessons1000/lesson13/student"
)

func main() {
	markA := student.Mark{Name: "A", Value: 100}
	markB := student.Mark{"B", 90}
	markC := student.Mark{"C", 80}
	markD := student.Mark{"D", 70}
	markE := student.Mark{"E", 60}
	markF := student.Mark{"F", 50}
	marks1 := []student.Mark{markA, markB, markC}
	st1 := &student.Student{
		FirstName: "yerassyl",
		LastName:  "tleugazy",
		Marks:     marks1,
	}
	st1.PrintFullInfo()
	marks2 := []student.Mark{
		markA, markB, markC, markD, markE, markF,
	}
	st2 := student.Student{
		FirstName: "asdasd",
		LastName:  "sdsd",
		Marks:     marks2,
	}
	st2.PrintFullInfo()
}
