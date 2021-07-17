package main

import "fmt"

func PrintData() {
	fmt.Println("Hello from print data")
}

func GetSum(a, b int) {
	fmt.Println(a + b)
}

func main() {
	PrintData()
	GetSum(6, 7)
	fmt.Println("Hello from main")
}
