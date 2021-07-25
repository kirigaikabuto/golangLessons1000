package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
	FullName  string
	Mark      float64
}

func main() {
	st1 := Student{
		FirstName: "1",
		LastName:  "2",
		FullName:  "",
		Mark:      3.6,
	}
	st1.FullName = st1.FirstName + " " + st1.LastName
	st1.FirstName = "3"
	st1.LastName = "4"
	fmt.Println(st1)
}
