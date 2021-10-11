package main

import "fmt"

func main() {
	var a, b int
	a = 3
	b = 4
	fmt.Println(a + b)
	arr := []int{1, 2, 3, 4, 45}
	arr = append(arr, 3)
	for i, v := range arr {
		fmt.Println(i, v)
	}
	d := map[string]string{
		"name":    "yerassyl",
		"surname": "tleugazy",
		"age":     "23",
	}
	d["address"] = "weqweqwqew"
	fmt.Println(d["name"])
	fmt.Println(d["surname"])
	fmt.Println(d["age"])
	for i, v := range d {
		fmt.Println(i, v)
	}
}
