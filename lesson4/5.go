package main

import "fmt"

func main() {
	//1.3 добавление элементов в slice
	//инициализация slice с заранее готовыми значениями
	//a1 := []int{1, 2, 3, 4, 5}
	//инизиализация пустого slice
	a2 := []int{}
	fmt.Println(a2)
	//code
	a2 = append(a2, 34)
	a2 = append(a2, 45)
	a2 = append(a2, 30)
	fmt.Println(a2)
	fmt.Println(a2[0])
}
