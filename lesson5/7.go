package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	newArr := []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] != 5 {
			newArr = append(newArr, arr[i])
		}
	}
	fmt.Println(newArr)
}
