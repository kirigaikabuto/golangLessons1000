package shop

import "fmt"

type Product struct {
	Name     string
	Price    int
	Category Category
}

func (p *Product) GetInfo() string {
	return fmt.Sprintf("Name:%s;Price:%d;Category:%s", p.Name, p.Price, p.Category.Name)
}

func (p *Product) PrintInfo() {
	fmt.Println(p.GetInfo())
}
