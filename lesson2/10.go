package main

import "fmt"

func main() {
	var salary int
	fmt.Scan(&salary)
	if salary >= 100 {
		fmt.Println("high")
	}
	if salary >= 50 && salary < 100 {
		fmt.Println("middle")
	}
	if salary < 50 {
		fmt.Println("low")
	}
}
