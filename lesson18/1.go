package main

import (
	"fmt"
	"github.com/kirigaikabuto/golanglessons1000/lesson18/products"
)

func main() {
	p1 := &products.Product{Name: "product1", Price: 300}
	p1.Save()
	fmt.Println(p1)
	productsElements := products.Product{}.GetAll()
	fmt.Println(productsElements)
}
