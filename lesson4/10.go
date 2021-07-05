package main

import "fmt"

func main() {
	//1.6 добавить элементы в slice используя клавиатуруу
	arr := []int{}
	n := 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		value := 0
		fmt.Scan(&value)
		arr = append(arr, value)
	}
	fmt.Println(arr)
}
