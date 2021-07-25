package main

import (
	"fmt"
	"math"
)

//https://acmp.ru/index.asp?main=task&id_task=26
func main() {
	var x1, y1, r1 float64
	var x2, y2, r2 float64
	fmt.Scan(&x1, &y1, &r1)
	fmt.Scan(&x2, &y2, &r2)
	part1 := math.Pow(x1-x2, 2)
	part2 := math.Pow(y1-y2, 2)
	distance := math.Sqrt(part1 + part2)
	if distance < r1+r2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
