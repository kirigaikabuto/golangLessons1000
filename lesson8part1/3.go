package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello my nam is Yerassyl"
	//split
	words := strings.Split(s, " ")
	for i := 0; i < len(words); i++ {
		//element Hello, my, name, is, Yerassyl
		element := words[i]
		for j := 0; j < len(element); j++ {
			//element=Hello, character -> H, e, l, l, o
			character := string(element[j])
			if character == "e" {
				fmt.Println(element)
				break
			}
		}
	}
}
