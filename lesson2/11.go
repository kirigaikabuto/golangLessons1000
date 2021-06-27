package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a <= 5 || a == 4 {
		fmt.Println("ты хорошо учишься")
	} else {
		fmt.Println("подтягивайся")
	}
}
