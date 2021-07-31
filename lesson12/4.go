package main

import "github.com/kirigaikabuto/golanglessons1000/lesson12/users"

func main() {
	a := users.Address{
		Name:            "1234",
		Number:          "123",
		HomePhoneNumber: "412312",
	}
	u1 := users.User{
		"1",
		"kirito",
		"123",
		a,
	}
	u1.PrintData()
}
