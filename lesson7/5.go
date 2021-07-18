package main

import "fmt"

func ConvertToOneDim(arr [][]int) []int {
	a := []int{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			a = append(a, arr[i][j])
		}
	}
	return a
}

func main() {
	test := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14},
	}
	fmt.Println(ConvertToOneDim(test))

}
