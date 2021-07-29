package main

import "fmt"

type User struct {
	Username string
}

func main() {
	user1 := &User{"1234"}
	user2 := user1
	fmt.Println(user1)
	fmt.Println(user2)
	user1.Username = "123"
	fmt.Println(user1)
	fmt.Println(user2)
}
