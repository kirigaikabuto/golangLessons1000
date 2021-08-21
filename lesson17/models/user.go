package models

import (
	"database/sql"
	"log"
)

type User struct {
	Id   string
	Name string
}

func InsertData(idValue, nameValue string, db *sql.DB) {
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
		log.Fatal(err)
		return
	}
}

func GetData(db *sql.DB) []User {
	users := []User{}
	selectQuery := "select id, name from users"
	rows, err := db.Query(selectQuery)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	for rows.Next() {
		item := User{}
		err = rows.Scan(&item.Id, &item.Name)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		users = append(users, item)
	}
	return users
}
