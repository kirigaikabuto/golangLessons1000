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
	item := todo.User{}
	err = json.Unmarshal(file, &item)
	if err != nil {
		fmt.Println(err)
		return
	}
	addressJson, err := json.Marshal(item.Address)
	if err != nil {
		fmt.Println(err)
		return
	}
	final := todo.UserInfo{}
	err = json.Unmarshal(file, &final)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(addressJson, &final)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileWrite, err := json.MarshalIndent(final, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile("final.json", fileWrite, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
