package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirigaikabuto/golanglessons1000/lesson15/todo"
	"io/ioutil"
)

func main() {
	user := todo.User{
		Username: "kirito",
		Address: todo.Address{
			Street: "zhevchenko",
			City:   "semey",
		},
	}
	file, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile("users.json", file, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
