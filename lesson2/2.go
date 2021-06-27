package main

import "fmt"

func main() {
	a := 3
	b := 4
	fmt.Println(a, b) //3 4
	a, b = b, a
	fmt.Println(a, b) //4 3
}
