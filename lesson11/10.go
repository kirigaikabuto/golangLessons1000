package main

import (
	"fmt"
	"github.com/kirigaikabuto/golanglessons1000/lesson11/users"
)

func main() {
	u1 := users.NewUser("kirito", "password")
	fmt.Println(u1.GetAllData())
}
