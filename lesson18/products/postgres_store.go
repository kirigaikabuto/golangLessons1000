package products

import (
	"database/sql"
	"fmt"
)

type productsPostgreStore struct {
	db *sql.DB
}

func NewPostgresProductStore() (ProductStore, error) {
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
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, err
	}
	return &productsPostgreStore{db: db}, nil
}

func (p *productsPostgreStore) CreateProduct(product Product) (*Product, error) {
	return nil, nil
}

func (p *productsPostgreStore) ListProduct() ([]Product, error) {
	return nil, nil
}

func (p *productsPostgreStore) GetProduct(id string) (*Product, error) {
	return nil, nil
}

func (p *productsPostgreStore) DeleteProduct(id string) error {
	return nil
}

func (p *productsPostgreStore) UpdateProduct(product Product) (*Product, error) {
	return nil, nil
}
