package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	createTableQuery := `
		create table if not exists users(
			id text,
			name text,
			primary key(id)
		);
	`
	host := "localhost"
	port := "5432"
	user := "setdatauser" //postgres
	database := "lesson17intro"
	password := "123456789" //admin
	params := "sslmode=disable"
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s", user, password, host, port, database, params)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
		return
	}
	//insert
	idValue := "4"
	nameValue := "kirito"
	queryInsert := "insert into users (id, name) values ($1, $2)"
	result, err := db.Exec(queryInsert, idValue, nameValue)
	if err != nil {
		log.Fatal(err)
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return
	}
	if n <= 0 {
		log.Fatal("some error during insert")
		return
	}

}
