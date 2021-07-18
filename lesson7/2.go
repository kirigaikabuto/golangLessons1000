package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{4, 5, 6, 7, 8}
	arrAll := [][]int{
		arr1,
		arr2,
		{4, 5, 6, 12, 23, 41, 3},
	}
	for i := 0; i < len(arrAll); i++ {
		for j := 0; j < len(arrAll[i]); j++ {
			fmt.Println(arrAll[i][j])
		}
	}
}
