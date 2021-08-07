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

func (s *Shop) GetCategories() []string {
	categories := []string{}
	for i := 0; i < s.GetCountOfProducts(); i++ {
		isExist := false
		for j := 0; j < len(categories); j++ {
			if s.Products[i].Category.Name == categories[j] {
				isExist = true
				break
			}
		}
		if !isExist {
			categories = append(categories, s.Products[i].Category.Name)
		}
	}
	return categories
}

func (s *Shop) PrintProductsCountByCategories() {
	categories := s.GetCategories()
	for j := 0; j < len(categories); j++ {
		count := 0
		for i := 0; i < s.GetCountOfProducts(); i++ {
			if s.Products[i].Category.Name == categories[j] {
				count += 1
			}
		}
		fmt.Println(fmt.Sprintf("%s:%d",categories[j],count))
	}
}
