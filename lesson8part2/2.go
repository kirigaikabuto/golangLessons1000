package main

import (
	"fmt"
	"strconv"
)

func main() {
	k := "45a"
	num, err := strconv.Atoi(k)
	if err != nil {
		fmt.Println("there is error:", err)
		return
	}
	fmt.Println(num, err)
	fmt.Println("end of program")
}
