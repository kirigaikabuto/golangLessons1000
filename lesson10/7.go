package main

import "fmt"

type Person1 struct {
	Name string
	Age  int
}

func main() {
	//необходимо создать структуру для человека
	//в ней должны быть поля name, age
	//создать три переменные подобного типа и заполнить значениями
	//вывести имя самого старшего
	p1 := Person1{
		Name: "",
		Age:  0,
	}
	p2 := Person1{
		Name: "",
		Age:  0,
	}
	myPersons := []Person1{p1, p2}

	persons := []Person1{
		{
			Name: "person1",
			Age:  12,
		},
		{
			Name: "person2",
			Age:  3,
		},
		{
			Name: "Person3",
			Age:  15,
		},
	}
	for i := 0; i < len(persons); i++ {
		fmt.Println(persons[i].Name, persons[i].Age)

	}
}
