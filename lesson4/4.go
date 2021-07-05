package main

import "fmt"

func main() {
	//1.2 заменить элемент slice
	a := []int{10, 11, 1, 2, 3}
	fmt.Println(a)
	a[1] = 4
	fmt.Println(a)
}
