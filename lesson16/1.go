package main

import "fmt"

func main() {
	person := map[string]int{
		"name":    123,
		"surname": 456,
	}
	fmt.Println(person["name"])
	fmt.Println(person["surname"])
}
