package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i]%2 == 0 {
			fmt.Println(arr[i])
		}
	}
	//найти сумму всех элементов массива и вывести ее
}
