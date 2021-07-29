package main

import "fmt"

type User struct {
	Username string
}

func changeUsername(u User) {
	u.Username = "111111"
}

func main() {
	user1 := User{"1234"}
	user2 := user1
	fmt.Println(user1)
	fmt.Println(user2)
	changeUsername(user1)
	fmt.Println(user1)
	fmt.Println(user2)
}
