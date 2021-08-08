package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kirigaikabuto/golanglessons1000/lesson15/todo"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://jsonplaceholder.typicode.com/todos"
	//create special body for todo
	todo1 := &todo.Todo{
		UserId:    3,
		Title:     "todo123",
		Completed: false,
	}
	//from struct to json
	jsonData, err := json.Marshal(todo1)
	if err != nil {
		fmt.Println(err)
		return
	}
	body := bytes.NewReader(jsonData)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json; charset=UTF-8'")
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
	//from json to struct
	item := &todo.Todo{}
	err = json.Unmarshal(dataJson, &item)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(item.Id)
	fmt.Println(item.UserId)
	fmt.Println(item.Completed)
	fmt.Println(item.Title)
}
