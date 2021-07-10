package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	//1>2 -> false
	//1>3 -> false
	//1>4 -> fale
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Printf("%d > %d -> %b \n", arr[i], arr[j], arr[i] > arr[j])
		}
	}
}
