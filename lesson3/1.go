package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	c := a % 2
	fmt.Println(c)
	if c == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Не четное")
	}
}
