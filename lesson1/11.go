package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3.0
	b := 3.0
	k := 16.0
	//math.Pow() -> возведение в степень
	c := math.Pow(a, b)
	fmt.Println(c)
	//math.Sqrt()-> квадратный корень
	d := math.Sqrt(k)
	fmt.Println(d)
}
