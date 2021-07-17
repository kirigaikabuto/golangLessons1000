package main

import "fmt"

func getMaxFromArr(a []int) (int, int) {
	maxi := 0
	maxiIndex := 0
	for i := 0; i < len(a); i++ {
		if a[i] > maxi {
			maxi = a[i]
			maxiIndex = i
		}
	}
	return maxi, maxiIndex
}

func main() {
	arr := []int{1, 2, 3, 9, 5, 6, 7, 8}
	m, i := getMaxFromArr(arr)
	fmt.Println(m, i)
}
