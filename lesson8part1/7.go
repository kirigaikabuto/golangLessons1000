package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GetString(s string) string {
	words := strings.Split(s, " ")
	result := ""
	//1) разделить слова и закинуть массив
	//2) создать пустую переменную строку
	//3) каждое слова из массива записывать в переменную строку и так же добавлять длину(склеивать)
	for i := 0; i < len(words); i++ {
		word := words[i]
		n := len(word)
		nString := strconv.Itoa(n)
		result += word + nString
	}
	return result
}

func main() {
	s := "Today is good day aSDSAD"
	s1 := "sdsdd sdsds dsds 123 sdasd"
	fmt.Println(GetString(s))
	fmt.Println(GetString(s1))
}
