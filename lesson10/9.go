package main

import "fmt"

type Book struct {
	Name  string
	Price int
}

func main() {
	b1 := Book{"book1", 300}
	b2 := Book{"book2", 400}
	books := []Book{
		b1,
		b2,
		{
			Name:  "book3",
			Price: 500,
		},
	}
	n := len(books)
	sumi := 0
	for i := 0; i < n; i++ {
		sumi = sumi + books[i].Price
	}
	avg := sumi / n
	for i := 0; i < n; i++ {
		if books[i].Price < avg {
			fmt.Println(books[i].Name)
		}
	}
	maxi := books[0]
	for i := 0; i < n; i++ {
		if maxi.Price < books[i].Price {
			maxi = books[i]
		}
	}
	fmt.Println(maxi)
}
