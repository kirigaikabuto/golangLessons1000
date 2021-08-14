package models

type Post struct {
	UserId float64 `json:"userId"`
	Id     float64`json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
