package main

import (
	"fmt"
	"github.com/kirigaikabuto/golanglessons1000/lesson13/products"
)

func main() {
	p1 := products.Product{
		Name:  "p1",
		Price: 100,
	}
	p2 := products.Product{
		Name:  "p2",
		Price: 200,
	}
	p3 := products.Product{
		Name:  "p3",
		Price: 300,
	}
	basket1 := products.Basket{
		Name:     "b1",
		Products: []products.Product{p1, p2},
	}
	fmt.Println("---Before add---")
	basket1.ListProducts()
	basket1.AddProduct(p3)
	fmt.Println("---After add---")
	basket1.ListProducts()
}
