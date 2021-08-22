package products

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
)

var (
	Database *sql.DB
)

type Product struct {
	Id    string
	Name  string
	Price int
}

//procedure which start when only start package
func init() {
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
	Database = db
}

func (p *Product) Save() {
	p.Id = uuid.New().String()
	queryInsert := "insert into products (id, name, price) values ($1, $2, $3)"
	result, err := Database.Exec(queryInsert, p.Id, p.Name, p.Price)
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

func (p Product) GetAll() []Product {
	products := []Product{}
	selectQuery := "select id, name, price from products"
	rows, err := Database.Query(selectQuery)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		item := Product{}
		err = rows.Scan(&item.Id, &item.Name, &item.Price)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		products = append(products, item)
	}
	return products
}
