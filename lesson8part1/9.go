package main

import (
	"fmt"
	"strings"
)

func ifSymbolNotExist(word, symbol string) bool {
	isNotExist := true
	for j := 0; j < len(word); j++ {
		if string(word[j]) == symbol {
			isNotExist = false
			break
		}
	}
	return isNotExist
}

func makeFilter(s, character string) string {
	result := ""
	words := strings.Split(s, " ")
	filtered := []string{}
	for i := 0; i < len(words); i++ {
		element := words[i]
		if ifSymbolNotExist(element, character) {
			filtered = append(filtered, element)
		}
	}
	result = strings.Join(filtered, " ")
	return result
}

func main() {
	s := "asdsd dfdfd qwetx xcxsdsew"
	fmt.Println(makeFilter(s, "a"))
}
