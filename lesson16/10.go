package main

import (
	"encoding/json"
	"fmt"
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
	fmt.Println(item["userId"], item["id"], item["title"], item["body"])
}
