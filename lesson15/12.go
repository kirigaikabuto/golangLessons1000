package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirigaikabuto/golanglessons1000/lesson15/todo"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("users.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	item := todo.UserFinal{}
	err = json.Unmarshal(file, &item)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileWrite, err := json.MarshalIndent(item, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile("users_final.json", fileWrite, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
