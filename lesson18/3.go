package main

import (
	"fmt"
	"github.com/kirigaikabuto/golanglessons1000/lesson18/products"
	"log"
)

func main() {
	productsStore, err := products.NewPostgresProductStore()
	if err != nil {
		log.Fatal(err)
		return
	}
	p1 := products.Product{
		Name:  "product3",
		Price: 1000,
	}
	newProduct, err := productsStore.CreateProduct(p1)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(newProduct)
}
