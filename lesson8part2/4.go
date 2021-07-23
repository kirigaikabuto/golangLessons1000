package main

import "fmt"

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main() {
	//сортировка (bubble sort, сортировка пурызьком)
	numbers := []int{23, 23, 2, 41, 9, -1, 4123}
	fmt.Println(numbers)
	//code
	numbers = bubbleSort(numbers)
	fmt.Println(numbers)
}
