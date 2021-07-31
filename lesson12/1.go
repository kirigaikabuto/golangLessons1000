package main

import "fmt"

func main() {
	c := 3
	var a *int
	a = &c
	fmt.Println(&c)
	fmt.Println(&a)

	fmt.Println(c)
	fmt.Println(a)

	fmt.Println(c)
	fmt.Println(*a)
}
