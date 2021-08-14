package main

import "fmt"

func main() {
	s := "how are 12333322"
	switch s {
	case "hello":
		fmt.Println("привет")
	case "how are you":
		fmt.Println("good")
	default:
		fmt.Println("error")
	}
}
