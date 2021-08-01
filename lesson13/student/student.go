package student

import (
	"fmt"
	"strings"
)

type Student struct {
	FirstName string
	LastName  string
	Marks     []Mark
}

func (s *Student) AddMark(mark Mark) {
	s.Marks = append(s.Marks, mark)
}

func (s *Student) PrintFullInfo() {
	res := "FirstName:%s; LastName:%s; Marks: [%s]; Avg: %d"
	marks := []string{}
	for i := 0; i < len(s.Marks); i++ {
		marks = append(marks, s.Marks[i].Name)
	}
	marksStr := strings.Join(marks, ",")
	avg := s.GetAverageMark()
	result := fmt.Sprintf(res, s.FirstName, s.LastName, marksStr, avg)
	fmt.Println(result)
}

func (s *Student) GetAverageMark() int {
	total := 0
	avg := 0
	for i := 0; i < len(s.Marks); i++ {
		total = total + s.Marks[i].Value
	}
	avg = total / len(s.Marks)
	return avg
}
