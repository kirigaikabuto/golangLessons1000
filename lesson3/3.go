package main

import "fmt"

func main() {
	var message string
	var n int
	fmt.Scanf("%s %d", &message, &n)
	for i := 0; i < n; i++ {
		fmt.Println(message)
	}
}
