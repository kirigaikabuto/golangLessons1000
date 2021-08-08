package todo

import "encoding/json"

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

type User struct {
	Username string  `json:"username"`
	Address  Address `json:"address"`
}

type UserInfo struct {
	Username string `json:"username"`
	Street   string `json:"street"`
	City     string `json:"city"`
}

type UserFinal struct {
	Username string `json:"username"`
	Street   string `json:"street"`
	City     string `json:"city"`
}

func (a *UserFinal) UnmarshalJSON(b []byte) error {
	var f interface{}
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	foomap := m["address"]
	v := foomap.(map[string]interface{})
	a.Username = m["username"].(string)
	a.Street = v["street"].(string)
	a.City = v["city"].(string)
	return nil
}
