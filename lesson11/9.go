package main

import (
	"fmt"
)

type Game struct {
	Name  string
	Price int
}

func (g Game) PrintData() {
	fmt.Println(g.Name)
}

func main() {
	g2 := Game{Name: "gavno"}
	fmt.Println(g2)
}
