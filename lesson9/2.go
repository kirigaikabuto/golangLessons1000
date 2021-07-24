package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateArray(e, l int) []int {
	arr := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= l; i++ {
		v := rand.Intn(e)
		arr = append(arr, v)
	}
	return arr
}

func generateTwoDimArray(n, m int) [][]int {
	res := [][]int{}
	for i := 0; i < n; i++ {
		arr := generateArray(100, m)
		res = append(res, arr)
	}
	return res
}

func main() {
	n := 3
	m := 4
	result := generateTwoDimArray(n, m)
	fmt.Println(result)
}
