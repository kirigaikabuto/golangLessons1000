package products

import "fmt"

type Product struct {
	Name  string
	Price int
}

func (p *Product) GetFullInfo() string {
	res := fmt.Sprintf("Name:%s,Price:%d", p.Name, p.Price)
	return res
}
