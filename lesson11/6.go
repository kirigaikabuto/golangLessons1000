package main

import "fmt"

type User6 struct {
	Username string
}

func (u *User6) changeName(name string) {
	u.Username = name
}

func (jkl User6) printName() {
	fmt.Println(jkl.Username)
}

func main() {
	u1 := &User6{"1234"}
	u1.printName()
	u1.changeName("4567")
	u1.printName()
	u2 := &User6{"Yerassyl"}
	u2.printName()
}
