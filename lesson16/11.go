package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirigaikabuto/golanglessons1000/lesson16/models"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://jsonplaceholder.typicode.com/posts/2"
	client := http.Client{}
	response, err := client.Get(url)
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
	item := map[string]interface{}{}
	err = json.Unmarshal(dataJson, &item)
	if err != nil {
		fmt.Println(err)
		return
	}
	post := &models.Post{}
	val, ok := item["userId"]
	if ok {
		post.UserId = val.(float64)
	}
	val, ok = item["id"]
	if ok {
		post.Id = val.(float64)
	}
	val, ok = item["title"]
	if ok {
		post.Title = val.(string)
	}
	val, ok = item["body"]
	if ok {
		post.Body = val.(string)
	}
	fmt.Println(post)
}
