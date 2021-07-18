package main

import "fmt"

func GetSumOfTwoDim(arr [][]int) []int {
	sums := []int{}
	for i := 0; i < len(arr); i++ {
		sumi := 0
		for j := 0; j < len(arr[i]); j++ {
			sumi += arr[i][j]
		}
		sums = append(sums, sumi)
	}
	return sums
}

func main() {
	test := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14},
	}
	fmt.Println(GetSumOfTwoDim(test))
}
