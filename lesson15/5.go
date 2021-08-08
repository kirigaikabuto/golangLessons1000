package main

import (
	"encoding/json"
	"fmt"
	todo2 "github.com/kirigaikabuto/golanglessons1000/lesson15/todo"
	"io/ioutil"
)

func main() {
	todo := todo2.Todo{
		UserId:    2,
		Id:        1,
		Title:     "example",
		Completed: false,
	}
	file, err := json.MarshalIndent(todo, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile("todo.json", file, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
