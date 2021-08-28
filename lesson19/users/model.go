package users

type User struct {
	Id       string `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type UserStore interface {
	Create(user *User) (*User, error)
	List() ([]User, error)
	Delete(id string) error
}
