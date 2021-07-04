package main

import "fmt"

func main() {
	for counter := 0; counter < 5; counter++ {
		fmt.Println(counter, "hello")
	}

	//1.counter = 0 -> Message counter+=1 counter<5 ->2
	//2.counter = 1 -> Message counter+=1 counter<5 ->3
	//3.counter = 2 -> Messsage

}
