package users

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type usersPostgreStore struct {
	db *sql.DB
}

func NewUsersPostgresStore(host, database, user, password, port string) (UserStore, error) {
	createTableQuery := `
		create table if not exists users(
			id text,
			full_name text,
			username text,
			password text,
			age int,
			primary key(id)
		);
	`
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
	return &usersPostgreStore{db: db}, nil
}

func (u *usersPostgreStore) Create(user *User) (*User, error) {
	user.Id = uuid.New().String()
	queryInsert := "insert into users (id, full_name, username, password, age) values ($1, $2, $3, $4, $5)"
	result, err := u.db.Exec(queryInsert, user.Id, user.FullName, user.Username, user.Password, user.Age)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, errors.New("error during creatring of user")
	}
	return user, nil
}

func (u *usersPostgreStore) List() ([]User, error) {
	users := []User{}
	selectQuery := "select id, full_name, username, password, age from users"
	rows, err := u.db.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		item := User{}
		err = rows.Scan(&item.Id, &item.FullName, &item.Username, &item.Password, &item.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, item)
	}
	return users, nil
}

func (u *usersPostgreStore) Delete(id string) error {
	result, err := u.db.Exec("delete from users where id= $1", id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n <= 0 {
		return errors.New("error in deleting")
	}
	return nil
}
