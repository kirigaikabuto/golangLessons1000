package main

import (
	"fmt"
	"github.com/kirigaikabuto/golanglessons1000/lesson19/users"
	"log"
)

func main() {
	host := "localhost"
	port := "5432"
	user := "setdatauser" //postgres
	database := "lesson18"
	password := "123456789" //admin
	store, err := users.NewUsersPostgresStore(host, database, user, password, port)
	if err != nil {
		log.Fatal(err)
		return
	}
	//u := &users.User{
	//	FullName: "Tleugazy Yerassyl",
	//	Username: "kirito",
	//	Password: "12345",
	//	Age:      23,
	//}
	//newUser, err := store.Create(u)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//fmt.Println(newUser)
	users, err := store.List()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(users)

}
