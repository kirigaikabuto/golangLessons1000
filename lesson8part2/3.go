package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(arr)
	arr[0], arr[1] = arr[1], arr[0]
	fmt.Println(arr)
}
