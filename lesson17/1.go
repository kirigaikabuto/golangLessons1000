package main

import (
	"fmt"
	"github.com/kirigaikabuto/golanglessons1000/lesson17/models"
	"log"
)

func main() {
	userStore := models.NewUserStore("localhost", "5432", "kirito", "12345")
	u := models.User{
		Username: "user1",
		Password: "user1",
	}
	newUser, err := userStore.Create(u)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(newUser)
}
