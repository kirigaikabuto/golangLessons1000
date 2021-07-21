package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"hello", "yerassyl", "19"}
	s := strings.Join(words, "+")
	fmt.Println(s)
}
