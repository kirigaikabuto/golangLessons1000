package users

func NewUser(username, password string) UserInter {
	return user{Username: username, Password: password}
}

type UserInter interface {
	GetUsername() string
	GetPassword() string
	GetAllData() string
}

type user struct {
	Username string
	Password string
}

func (u user) GetUsername() string {
	return u.Username
}

func (u user) GetPassword() string {
	return u.Password
}

func (u user) GetAllData() string {
	return u.Username + " " + u.Password
}
