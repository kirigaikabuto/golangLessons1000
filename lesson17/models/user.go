package models

import (
	"database/sql"
	"log"
)

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
