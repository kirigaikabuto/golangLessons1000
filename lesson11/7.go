package main

import "fmt"

type Book struct {
	Name   string
	Price  float64
	Author string
}

func (b Book) printData() {
	fmt.Printf("Name:%s,Price:%f,Author:%s \n", b.Name, b.Price, b.Author)
}

func (b *Book) setAuthorName(name string) {
	b.Author = name
}

func (b Book) getAuthorName() string {
	return b.Author
}

func main() {
	b1 := Book{
		Name:   "book1",
		Price:  34.5,
		Author: "yerassyl",
	}
	b1.printData()
	b2 := Book{
		Name:   "book2",
		Price:  100.5,
		Author: "yerassyl",
	}
	b2.printData()
	b1.setAuthorName("1234")
	b2.setAuthorName("4566")
	b1.printData()
	b2.printData()
	fmt.Println(b1.getAuthorName())
	fmt.Println(b2.getAuthorName())
}
