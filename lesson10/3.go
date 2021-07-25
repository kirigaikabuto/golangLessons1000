package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	//возраст и имя человека
	var p Person
	p.Name = "Yerassyl"
	p.Age = 23
	fmt.Println(p)
	p2 := Person{"sdsdsd", 15}
	fmt.Println(p2)
	p3 := Person{Name: "12314"}
	fmt.Println(p3)
}
