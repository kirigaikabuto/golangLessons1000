package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	for i, v := range arr {
		fmt.Println(i,v)
	}
}
