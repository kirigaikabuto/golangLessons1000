package main

import "fmt"

type Student2 struct {
	Username string
	Marks    []int
}

func main() {
	marks := []int{1, 23, 4, 5, 3, 2}
	st1 := Student2{
		Username: "123",
		Marks:    []int{1, 2, 3, 4, 5, 6},
	}
	st2 := Student2{
		Username: "456",
		Marks:    marks,
	}
	n := len(st1.Marks)
	sumi := 0
	for i := 0; i < n; i++ {
		sumi += st1.Marks[i]
	}
	fmt.Println(sumi)
	fmt.Println(st1)
	fmt.Println(st2)
}
