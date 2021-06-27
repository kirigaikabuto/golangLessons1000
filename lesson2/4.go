package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	c := a + b
	fmt.Println(c)
}
