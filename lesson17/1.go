package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	//createTableQuery := `
	//	create table if not exits users(
	//		id text,
	//		name text,
	//		primary key(id)
	//	);
	//`
	host := "localhost"
	port := "5432"
	user := "setdatauser"
	database := "lesson17intro"
	password := "123456789"
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
}
