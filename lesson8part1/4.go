package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "1,2,3,4,5,7"
	words := strings.Split(s, ",")
	fmt.Println(words[0])
}
