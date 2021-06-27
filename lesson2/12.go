package main

import "fmt"

func main() {
	var salary int
	fmt.Scan(&salary)
	if salary >= 100 {
		fmt.Println("high")
	} else if salary >= 50 && salary < 100 {
		fmt.Println("middle")
	} else if salary < 50 {
		fmt.Println("low")
	} else {
		fmt.Println("errrr")
	}
}
