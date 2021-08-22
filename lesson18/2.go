package main

import (
	"database/sql"
	"fmt"
	"github.com/kirigaikabuto/golanglessons1000/lesson18/products"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"strings"
)

func main() {
	//connection to database
	createTableQuery := `
		create table if not exists products(
			id text,
			name text,
			price int,
			primary key(id)
		);
	`
	host := "localhost"
	port := "5432"
	user := "setdatauser" //postgres
	database := "lesson18"
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
	product := &products.Product{
		Id:    "f314e547-aba1-48cc-9026-606023862f2b",
		Name:  "1231232",
		Price: 1000,
	}
	queryUpdate := "update products set "
	parts := []string{}
	values := []interface{}{}
	i := 0
	if product.Name != "" {
		i++
		parts = append(parts, "name=$"+strconv.Itoa(i)) //name=$1
		values = append(values, product.Name)
	}
	if product.Price != 0 {
		i++
		parts = append(parts, "price=$"+strconv.Itoa(i)) //name=$1
		values = append(values, product.Price)
	}
	i++
	queryUpdate = queryUpdate + strings.Join(parts, ",") + " where id=$" + strconv.Itoa(i)
	fmt.Println(queryUpdate)
	values = append(values, product.Id)
	result, err := db.Exec(queryUpdate, values...)
	if err != nil {
		log.Fatal(err)
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
