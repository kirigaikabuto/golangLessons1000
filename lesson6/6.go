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

func main() {
	k := generateArray(10, 5)
	fmt.Println(k)

}
