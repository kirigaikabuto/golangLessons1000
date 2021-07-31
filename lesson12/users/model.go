package users

import "fmt"

type User struct {
	Id       string
	Username string
	Password string
}

func (u *User) PrintData() {
	fmt.Println(u.Id, u.Username, u.Password)
}
