package main

import "fmt"

type Man struct {
	Name string
	Age  int
}

func main() {

	m1 := Man{
		Name: "Born",
		Age:  45,
	}
	m2 := Man{
		Name: "Stetcham",
		Age:  46,
	}
	m3 := Man{
		Name: "Artur",
		Age:  50,
	}

	maxi := m1.Age
	person := m1.Name
	if maxi <= m2.Age {
		maxi = m2.Age
		person = m2.Name
	} else if maxi <= m3.Age {
		maxi = m3.Age
		person = m3.Name
	} else {
		maxi = m1.Age
		person = m1.Name
	}

	fmt.Println(maxi, person)

}
