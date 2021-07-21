package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Today is good day aSDSAD"
	//1) превратить строку в массив из слов
	//2) пробежаться по массиву и запоминать длину каждого слова
	//3) определять слово с максимальной длинной
	words := strings.Split(s, " ")
	maxiLength := len(words[0])
	maxiWord := words[0]
	for i := 0; i < len(words); i++ {
		if maxiLength < len(words[i]) {
			maxiLength = len(words[i])
			maxiWord = words[i]
		}
	}
	fmt.Println(maxiLength, maxiWord)
}
