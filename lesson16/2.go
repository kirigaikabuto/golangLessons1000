package main

import "fmt"

func main() {
	d := map[string]string{}
	fmt.Println(d)
	d["name"] = "yerassyl"
	d["age"] = "23"
	fmt.Println(len(d))
	for i, v := range d {
		fmt.Println(i, v)
	}
}
