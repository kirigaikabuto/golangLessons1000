package shop

import (
	"fmt"
	"strings"
)

type Shop struct {
	Name     string
	Products []Product
}

func (s *Shop) PrintProducts(header string) {
	fmt.Println(fmt.Sprintf("\t---%s---", strings.ToUpper(header)))
	for i := 0; i < len(s.Products); i++ {
		s.Products[i].PrintInfo()
	}
}

func (s *Shop) GetAveragePrice() int {
	n := len(s.Products)
	sumi := 0
	avg := 0
	for i := 0; i < n; i++ {
		sumi += s.Products[i].Price
	}
	avg = sumi / n
	return avg
}

func (s *Shop) GetCountOfProducts() int {
	n := len(s.Products)
	return n
}

func (s *Shop) AddProduct(product Product) {
	s.Products = append(s.Products, product)
}

func (s *Shop) RemoveProduct(id string) {
	products := []Product{}
	//
	for i := 0; i < s.GetCountOfProducts(); i++ {
		if s.Products[i].Id != id {
			products = append(products, s.Products[i])
		}
	}
	s.Products = products
}

func (s *Shop) BuyProduct(id string, count int) {
	for i := 0; i < s.GetCountOfProducts(); i++ {
		if s.Products[i].Id == id {
			s.Products[i].Count = s.Products[i].Count - count
		}
	}
}
