package main

import "fmt"

func main() {
	//1.1 получить значени из slice
	//slice
	b := []int{11, 32, 34, 45, 5, 6, 7, 8}
	//indexs
	//values: 11 32 34 45 5 6 7 8
	//indexs: 0  1  2  3  4 5 6 7
	fmt.Println(b)
	first := b[0]
	fmt.Println("первый элемент массива", first)
	fmt.Println("пятый элемент массива", b[4])
	fmt.Println("сумма элементов массива", b[1]+b[5])
}
