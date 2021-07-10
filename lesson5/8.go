package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	arr1 := arr[:4]
	fmt.Println(arr)
	fmt.Println(arr1)
	//[start_index:end_index]
	arr2 := arr[4:]
	fmt.Println(arr2)
}
