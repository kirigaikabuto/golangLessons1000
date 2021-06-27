package main

import "fmt"

func main() {
	a := 3
	b := 4
	fmt.Println(a, b) //3 4
	c := a
	a = b
	b = c
	fmt.Println(a, b) //4 3
}
