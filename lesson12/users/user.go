package users

import "fmt"

type User struct {
	Id          string
	Username    string
	Password    string
	HomeAddress Address
}

func (u *User) PrintData() {
	fmt.Println(u.Id, u.Username, u.Password, u.HomeAddress.GetData())
}
