package main

import "fmt"

func main() {
	//break
	//continue
	n := 10
	fmt.Println("start cycle")
	for i := 0; i < n; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("end cycle")
}
