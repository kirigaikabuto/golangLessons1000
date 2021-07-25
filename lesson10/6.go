package main

import "fmt"

type Student3 struct {
	Name  string
	Marks []int
}

func main() {
	st1 := Student3{
		Name:  "student1",
		Marks: []int{1, 2, 3, 4, 5},
	}
	fmt.Println(st1)
	st1.Marks[0] = 125
	fmt.Println(st1)
}
