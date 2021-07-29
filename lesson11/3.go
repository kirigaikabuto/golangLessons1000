package main

import "fmt"

func main() {
	var a int
	a = 3
	var b *int
	b = &a
	fmt.Println(a)
	fmt.Println(*b)
	a = 4
	fmt.Println(a)
	fmt.Println(*b)
	*b = 120
	fmt.Println(a)
	fmt.Println(*b)
}
