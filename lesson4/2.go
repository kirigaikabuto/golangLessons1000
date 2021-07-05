package main

import "fmt"

func main() {
	n := 10
	fmt.Println("start cycle")
	for i := 0; i < n; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("end cycle")
}
