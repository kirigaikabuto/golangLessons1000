package main

import "fmt"

func main() {
	var c interface{}
	c = false
	switch c.(type) {
	case int:
		fmt.Println("это число")
	case string:
		fmt.Println("это строка")
	case bool:
		fmt.Println("логический тип")
	default:
		fmt.Println("неизствестно")
	}
}
