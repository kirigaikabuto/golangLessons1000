package main

import "fmt"

func main() {
	//int,float64,string,boolean
	var c bool
	a := 3
	b := 4
	//+-/*
	//> < >= <= == !=
	c = a < b
	c = a != b
	c = a == b
	fmt.Println(c)
}
