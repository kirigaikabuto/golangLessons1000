package main

import "fmt"

type User struct {
	Name     string
	Surname  string
	Age      int
	FullName string
}

func (u *User) ShowAllInfo() {
	fmt.Println(u.Name, u.Surname, u.Age, u.FullName)
}

func (u *User) SetFullName() {
	u.FullName = u.Name + " " + u.Surname
}
func main() {
	u1 := &User{
		Name:    "Yerassyl",
		Surname: "Tleugazy",
		Age:     23,
	}
	u1.ShowAllInfo()
	u1.SetFullName()
	u1.ShowAllInfo()
}
