package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{4, 5, 6, 7, 8}
	arrAll := [][]int{
		arr1,
		arr2,
		{4,5,6,12,23,41,3,},
	}
	fmt.Println(arrAll[1][4])
	//rows, columns
	//  0 1 2 3 4 5
	//0 1 2 3 4 5
	//1 4 5 6 7 8
	//2
	//3
	//4
	//5
}
