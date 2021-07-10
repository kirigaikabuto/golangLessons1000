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
	i := 0
	sumi := 0
	for i = 0; i < n; i++ {
		k := rand.Intn(limit)
		arr = append(arr, k)
	}
	for i = 0; i < n; i++ {
		sumi = sumi + arr[i]
	}
	fmt.Println(arr)
	fmt.Println(sumi / n)
}
