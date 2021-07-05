package main

import "fmt"

func main() {
	//1.5 как вывести все элементы массива используя цикл
	a := []int{1, 2, 3, 4, 2, 23, 12, 32}
	for i := 0; i < len(a); i++ {
		fmt.Printf("index:%d;value:%d \n", i, a[i])
	}
}
