package main

import "fmt"

func bubbleSort2(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func getArrayOfSum(numbers [][]int) []int {
	sums := []int{}
	for i := 0; i < len(numbers); i++ {
		sumi := 0
		for j := 0; j < len(numbers[i]); j++ {
			sumi += numbers[i][j]
		}
		sums = append(sums, sumi)
	}
	return sums
}

func main() {
	numbers := [][]int{
		{3, 1, 23, 541},
		{1, 23, 412, 23},
		{4, 2, 23, 42, 13, 42},
	}
	fmt.Println(numbers)
	//code
	for i := 0; i < len(numbers); i++ {
		numbers[i] = bubbleSort2(numbers[i])
	}
	fmt.Println(numbers)
	sums := getArrayOfSum(numbers)
	fmt.Println(sums)

}
