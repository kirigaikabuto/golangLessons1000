package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	a := 5
	d := float64(a)
	c := int(d)
	res := math.Pow(d, 2.0)
	fmt.Println(res)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(res))
}
