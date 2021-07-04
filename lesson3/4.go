package main

import "fmt"

func main() {
	//0 + 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9
	sumi := 0
	n := 10
	//sumi = sumi + 0 //0
	//sumi = sumi + 1 //1
	//sumi = sumi + 2 //3
	//sumi = sumi + 3 //6
	//sumi = sumi + 4 //10
	//sumi = sumi + 5 //15
	//sumi = sumi + 6 //21
	//sumi = sumi + 7 //28
	for i := 0; i < n; i++ {
		sumi = sumi + i
		fmt.Println(sumi)
	}
	fmt.Println(sumi)
}
