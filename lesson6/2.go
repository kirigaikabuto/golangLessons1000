package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	arr1 := arr[:4]
	arr2 := arr[4:]
	fmt.Println(arr1)
	fmt.Println(arr2)
}
