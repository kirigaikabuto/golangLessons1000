package main

import "fmt"

func main() {
	//1.4 длина slice
	a := []int{}
	n := len(a)
	fmt.Println(a)
	fmt.Println(n)
	a = append(a, 3)
	fmt.Println(a)
	fmt.Println(len(a))
}
