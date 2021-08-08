package main

import (
	"encoding/json"
	"fmt"
	todo "github.com/kirigaikabuto/golanglessons1000/lesson15/todo"
	"io/ioutil"
)

func main() {
	todo1 := todo.Todo{
		UserId:    2,
		Id:        1,
		Title:     "example",
		Completed: false,
	}
	todo2 := todo.Todo{
		UserId:    3,
		Id:        2,
		Title:     "asdsds",
		Completed: false,
	}
	todos := []todo.Todo{todo1, todo2}
	file, err := json.MarshalIndent(todos, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile("todos.json", file, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
