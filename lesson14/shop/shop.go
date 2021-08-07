package shop

type Shop struct {
	Name     string
	Products []Product
}

func (s *Shop) PrintProducts() {
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
