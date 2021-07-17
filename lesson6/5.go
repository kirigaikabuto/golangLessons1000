package main

import "fmt"

func GetSumOfArr(a []int) int {
	sumi := 0
	for i := 0; i < len(a); i++ {
		sumi = sumi + a[i]
	}
	return sumi
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 2}
	arr2 := []int{33, 41, 123, 1}
	k1 := GetSumOfArr(arr1)
	k2 := GetSumOfArr(arr2)
	fmt.Println(k1)
	fmt.Println(k2)
}
