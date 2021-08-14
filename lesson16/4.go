package main

import "fmt"

func main() {
	d := map[string]string{}
	d["name"] = "yerassyl"
	d["age"] = "23"
	for i, v := range d {
		fmt.Println(i, v)
	}
}
