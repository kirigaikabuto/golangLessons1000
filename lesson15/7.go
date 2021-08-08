package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirigaikabuto/golanglessons1000/lesson15/todo"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("todo.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	item := todo.Todo{}
	err = json.Unmarshal(file, &item)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(item.Id)
	fmt.Println(item.Title)
	fmt.Println(item.Completed)
	fmt.Println(item.UserId)
}
