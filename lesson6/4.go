package main

import (
	"fmt"
)

func GetSumOfNumbers(a, b int) int {
	c := a + b
	return c
}

func main() {
	k := GetSumOfNumbers(3, 4)
	l := GetSumOfNumbers(5, 6)
	fmt.Println(k)
	fmt.Println(l)
}
