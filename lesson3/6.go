package main

import "fmt"

func main() {
	s := 0
	e := 0
	fmt.Scanf("%d %d", &s, &e)
	sum := 0
	mlt := 1
	n := e - s
	half := s + n/2
	fmt.Println(half)
	for i := s; i < half; i++ {
		sum = sum + i
	}
	for i := half; i < e; i++ {
		mlt = mlt * i
	}
	fmt.Println("сумма=", sum, "умножение=", mlt)
}
