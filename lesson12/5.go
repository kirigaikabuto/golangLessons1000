package main

import "github.com/kirigaikabuto/golanglessons1000/lesson12/users"

func main() {
	u1 := users.User{
		Id:       "2",
		Username: "yerassyl",
		Password: "1234",
		HomeAddress: users.Address{
			Name:            "addess_anem",
			Number:          "13",
			HomePhoneNumber: "123123113",
		},
	}
	u1.PrintData()
}
