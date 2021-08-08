package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirigaikabuto/golanglessons1000/lesson15/todo"
	"io/ioutil"
	"net/http"
)

func main() {
	//get many elements
	url := "http://jsonplaceholder.typicode.com/todos"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json")
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	dataJson, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	items := []todo.Todo{}
	err = json.Unmarshal(dataJson, &items)
	if err != nil {
		fmt.Println(err)
		return
	}
	//write to file
	file, err := json.MarshalIndent(items, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile("todos_from_api.json", file, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
