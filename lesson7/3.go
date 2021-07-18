package main

import "fmt"

func getMaxForTwoDim(arr [][]int) int {
	maxi := arr[0][0]
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if maxi < arr[i][j] {
				maxi = arr[i][j]
			}
		}
	}
	return maxi
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{4, 5, 6, 7, 8}
	arrAll := [][]int{
		arr1,
		arr2,
		{4, 5, 6, 12, 23, 41, 3},
	}
	fmt.Println(getMaxForTwoDim(arrAll))

}
