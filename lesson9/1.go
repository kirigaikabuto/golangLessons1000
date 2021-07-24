package main

import "fmt"

func RemoveFromArray(arr []int, element int) []int {
	newArr := []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] != element {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 4}
	elementToDelete := 4
	fmt.Println(arr)
	arr = RemoveFromArray(arr, elementToDelete)
	fmt.Println(arr)
}
