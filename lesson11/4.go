package main

import "fmt"

func addNumber(a *int) {
	*a = *a + 1
}

func main() {
	l := 3
	//var c *int
	//c = &l
	fmt.Println(l)
	addNumber(&l)
	fmt.Println(l)
}
