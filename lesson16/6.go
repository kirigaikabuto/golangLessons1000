package main

import (
	"fmt"
	"reflect"
)

func main() {
	var c int64 = 3
	var a interface{}
	a = c
	fmt.Println(a, reflect.TypeOf(a))
}
