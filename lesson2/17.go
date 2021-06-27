package main

import "fmt"

func main() {
	var znak string
	var a, b int
	fmt.Scanf("%s %d %d", &znak, &a, &b)
	res := 0
	if znak == "+" {
		res = a + b
	} else if znak == "-" {
		res = a - b
	} else if znak == "*" {
		res = a * b
	}
	fmt.Printf("%d%s%d=%d", a, znak, b, res)

	//+ 3 4
	//3+4=7
	//* 4 5
	//4*5=20
}
