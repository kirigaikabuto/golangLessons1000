package users

import (
	_ "github.com/lib/pq"
)

type usersMongoStore struct {
}

func NewUsersMongoStore() (UserStore, error) {

	return &usersMongoStore{}, nil
}

func (u *usersMongoStore) Create(user *User) (*User, error) {
	return nil, nil
}

func (u *usersMongoStore) List() ([]User, error) {
	return nil, nil
}

func (u *usersMongoStore) Delete(id string) error {
	return nil
}
