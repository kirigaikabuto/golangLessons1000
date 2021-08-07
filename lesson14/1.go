package main

import (
	"fmt"
	"github.com/kirigaikabuto/golanglessons1000/lesson14/shop"
)

func main() {
	category1 := shop.Category{Name: "mobile"}
	product1 := shop.Product{
		Name:     "samsung",
		Price:    1,
		Category: category1,
	}
	product2 := shop.Product{
		Name:     "iphone",
		Price:    2,
		Category: category1,
	}
	category2 := shop.Category{Name: "pc"}
	product3 := shop.Product{
		Name:     "legion",
		Price:    3,
		Category: category2,
	}
	product4 := shop.Product{
		Name:     "hp omen",
		Price:    4,
		Category: category2,
	}
	shopAlser := shop.Shop{
		Name:     "alser",
		Products: []shop.Product{product1, product2, product3, product4},
	}
	shopAlser.PrintProducts()
	fmt.Println(shopAlser.GetAveragePrice())
}
