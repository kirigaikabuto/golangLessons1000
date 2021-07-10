package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//использовать библиотеку rand чтобы генерировать массив из чисел
	rand.Seed(time.Now().UnixNano())
	min := 50
	max := 100
	k := rand.Intn(max-min) + min
	l := rand.Intn(max-min) + min
	fmt.Println(k, l)
}
