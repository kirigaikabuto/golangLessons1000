package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"hello", "yerassyl"}
	s := strings.Join(words, " ")
	fmt.Println(s)
}
