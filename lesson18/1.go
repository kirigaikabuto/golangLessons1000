package main

import "github.com/kirigaikabuto/golanglessons1000/lesson18/products"

func main() {
	//p1 := &products.Product{Name: "product1", Price: 300}
	//p1.Save()
	//fmt.Println(p1)
	//productsElements := products.Product{}.GetAll()
	//fmt.Println(productsElements)
	//
	p2 := &products.Product{Id: "f314e547-aba1-48cc-9026-606023862f2b", Name: "11111"}
	p2.Update()
}
