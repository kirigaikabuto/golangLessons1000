package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	arr = append(arr, 3)
	newArr := [100]int{}
	fmt.Println(cap(newArr))
	fmt.Println(cap(arr))
	fmt.Println(newArr)
}
