package main

import "fmt"

type Book struct {
	Name  string
	Price float64
}

func (b *Book) PrintData() {
	b.setDefault()
	fmt.Println(b.Name, b.Price)
}

func (b *Book) setDefault() {
	if b.Name == "" {
		b.Name = "default name"
	}
	if b.Price == 0 {
		b.Price = 3.4
	}
}

func main() {
	b1 := Book{"asdaads",10.5}
	b1.PrintData()
}
