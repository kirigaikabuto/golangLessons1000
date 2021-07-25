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
	example := Person1{}
	for i := 0; i < len(persons); i++ {
		fmt.Println(persons[i].Name, persons[i].Age)
		example = persons[i]
	}
}
