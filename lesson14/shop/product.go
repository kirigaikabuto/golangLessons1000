package shop

import "fmt"

type Product struct {
	Id       string
	Name     string
	Price    int
	Count    int
	Category Category
}

func (p *Product) GetInfo() string {
	return fmt.Sprintf("Id:%s,Name:%s;Price:%d;Count:%d;Category:%s", p.Id, p.Name, p.Price, p.Count, p.Category.Name)
}

func (p *Product) PrintInfo() {
	fmt.Println(p.GetInfo())
}
