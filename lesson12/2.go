package main

import "fmt"

type User struct {
	Username  string
	Password  string
	FirstName string
	LastName  string
	FullName  string
}

func (u User) printData() {
	u.setFullName()
	fmt.Println(u.Username, u.Password, u.FirstName, u.LastName, u.FullName)
}

func (u *User) setFullName() {
	u.FullName = u.FirstName + u.LastName
}

func main() {
	u1 := User{
		Username:  "user1",
		Password:  "user1",
		FirstName: "user1",
		LastName:  "user1",
	}
	u1.printData()
	u2 := User{
		Username:  "user2",
		Password:  "user2",
		FirstName: "user2",
		LastName:  "user2",
	}
	u2.printData()
}
