package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	limit := 100
	arr := []int{}
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		k := rand.Intn(limit)
		arr = append(arr, k)
	}
	maxi := arr[0]
	maxiIndex := 0
	for i := 0; i < n; i++ {
		if arr[i] > maxi {
			maxi = arr[i]
			maxiIndex = i
		}
	}
	mini := arr[0]
	miniIndex := 0
	for i := 0; i < n; i++ {
		if arr[i] < mini {
			mini = arr[i]
			miniIndex = i
		}
	}
	fmt.Println(arr)
	c := arr[maxiIndex]
	arr[maxiIndex] = arr[miniIndex]
	arr[miniIndex] = c
	fmt.Println(arr)
}
