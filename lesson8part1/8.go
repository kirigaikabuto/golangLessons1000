package main

import (
	"fmt"
	"strings"
)

func filterString(s string) string {
	result := ""
	//code
	words1 := strings.Split(s, " ")
	words2 := []string{}
	for i := 0; i < len(words1); i++ {
		if words1[i] != "or" && words1[i] != "and" {
			words2 = append(words2, words1[i])
		}
	}
	result = strings.Join(words2, " ")
	return result
}

func main() {
	s := "get set or allll check and"
	fmt.Println(filterString(s))
}
