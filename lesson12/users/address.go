package users

import "fmt"

type Address struct {
	Name            string
	Number          string
	HomePhoneNumber string
}

func (a *Address) PrintData() {
	fmt.Println(a.Name, a.Number, a.HomePhoneNumber)
}

func (a *Address) GetData() string {
	s := fmt.Sprintf("HomeName:%s, HomeNumber:%s, HomePhoneNumber:%s", a.Name, a.Number, a.HomePhoneNumber)
	return s
}
