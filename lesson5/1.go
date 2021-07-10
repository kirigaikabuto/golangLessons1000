package main

import "fmt"

func main() {
	arr := []int{100, 22, 33, 44, 5, 6}
	//1) необходимо создать переменную в которой будет храниться максимальное значение
	//
	maxi := arr[0]
	for i := 0; i < len(arr); i++ {
		//2) написать условие для нахождения максимального числа
		if arr[i] > maxi {
			maxi = arr[i]
		}
	}
	fmt.Println(maxi)
}
