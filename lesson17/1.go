package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func insertData(idValue, nameValue string, db *sql.DB) error {
	queryInsert := "insert into users (id, name) values ($1, $2)"
	result, err := db.Exec(queryInsert, idValue, nameValue)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n <= 0 {
		return errors.New("error during creating of user")
	}
	return nil
}

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
	err = insertData("5", "yerassl", db)
	if err != nil {
		log.Fatal(err)
		return
	}

}
