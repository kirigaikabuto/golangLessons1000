package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := len(arr)
	sum := 0
	l := 0
	for i := 0; i < n; i++ {
		l = 3
		fmt.Println(l)
		sum = sum + arr[i]
	}
	fmt.Println(l)
	fmt.Println("сумма элементов массива =", sum)
}
