package main

import "fmt"

func main() {
	//var s string
	//s = "Hello"
	var c string
	fmt.Scan(&c)
	if c == "Yerassyl" {
		fmt.Println("Hello Yerassyl")
	} else {
		fmt.Println("unknown")
	}
}
