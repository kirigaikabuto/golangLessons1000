package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirigaikabuto/golanglessons1000/lesson15/todo"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("todos.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	items := []todo.Todo{}
	err = json.Unmarshal(file, &items)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(items); i++ {
		fmt.Println(items[i].Id, items[i].Title)
	}
}
