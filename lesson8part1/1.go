package main

import "fmt"

func main() {
	arr := []string{"Hello", "Yerassyl", "lesson"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}

}
