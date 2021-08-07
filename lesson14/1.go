package main

import (
	"fmt"
	"github.com/kirigaikabuto/golanglessons1000/lesson14/shop"
)

func main() {
	category1 := shop.Category{Name: "mobile"}
	product1 := shop.Product{
		Id:       "1",
		Name:     "samsung",
		Price:    1,
		Category: category1,
		Count:    5,
	}
	product2 := shop.Product{
		Id:       "2",
		Name:     "iphone",
		Price:    2,
		Category: category1,
		Count:    5,
	}
	category2 := shop.Category{Name: "pc"}
	product3 := shop.Product{
		Id:       "3",
		Name:     "legion",
		Price:    3,
		Category: category2,
		Count:    5,
	}
	product4 := shop.Product{
		Id:       "4",
		Name:     "hp omen",
		Price:    4,
		Category: category2,
		Count:    5,
	}
	shopAlser := shop.Shop{
		Name:     "alser",
		Products: []shop.Product{product1, product2, product3, product4},
	}
	category3 := shop.Category{Name: "tv"}
	product5 := shop.Product{
		Id:       "5",
		Name:     "lg",
		Price:    6,
		Category: category3,
		Count:    5,
	}
	shopAlser.AddProduct(product5)
	shopAlser.PrintProducts("add product result")
	shopAlser.BuyProduct("3", 4)
	shopAlser.PrintProducts("after buying of product by id 4")
	fmt.Println(shopAlser.GetCategories())
	shopAlser.PrintProductsCountByCategories()
	//fmt.Println(shopAlser.GetAveragePrice())
	//fmt.Println(shopAlser.GetCountOfProducts())
	//shopAlser.RemoveProduct("3")
	//shopAlser.PrintProducts()
}
