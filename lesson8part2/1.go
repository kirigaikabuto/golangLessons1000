package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getWords(pred string) []string {
	preds := strings.Split(pred, ".")
	words := []string{}
	for i := 0; i < len(preds); i++ {
		someWords := strings.Split(preds[i], " ")
		words = append(words, someWords...)
	}
	return words
}

func getNumbersAndWords(words []string) ([]int, []string) {
	numbers := []int{}
	others := []string{}
	for i := 0; i < len(words); i++ {
		word := words[i]
		num, err := strconv.Atoi(word)
		if err != nil {
			others = append(others, word)
			continue
		}
		numbers = append(numbers, num)
	}
	return numbers, others
}

func main() {
	//map dictionary словарь
	pred := "Hello my names is Yerassyl.I am living in Almaty 8903.Good day today 1994"
	words := getWords(pred)
	numbers, others := getNumbersAndWords(words)
	fmt.Println(numbers)
	fmt.Println(others)
}
