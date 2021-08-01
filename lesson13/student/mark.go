package student

import "fmt"

type Mark struct {
	Name  string
	Value int
}

func (m *Mark) GetInfo() string {
	res := fmt.Sprintf("Name:%s;Value:%d", m.Name, m.Value)
	return res
}
