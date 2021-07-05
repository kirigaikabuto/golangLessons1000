package main

import "fmt"

func main() {
	c := []int{}
	a := []int{}
	b := []int{}
	n := 20
	for i := 1; i <= n; i++ {
		c = append(c, i)
	}
	fmt.Println(c)
	for i := 0; i < len(c); i++ {
		if c[i]%2 == 0 {
			a = append(a, c[i])
		} else {
			b = append(b, c[i])
		}
	}
	fmt.Println(a)
	fmt.Println(b)

}
