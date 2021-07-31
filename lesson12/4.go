package main

import "github.com/kirigaikabuto/golanglessons1000/lesson12/users"

func main() {
	u1 := users.User{
		"123",
		"kirito",
		"123",
	}
	u1.PrintData()
}
