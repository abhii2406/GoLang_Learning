package user

import (
	"fmt"
	"time"
)

type User struct {
	name    string
	age     int
	curDate time.Time
}

func NewUser(name string, age int, curDate time.Time) *User {
	return &User{
		name:    name,
		age:     age,
		curDate: curDate,
	}
}
func (u User) PrintAll() {
	fmt.Println(u.name, " ", u.age, " ", u.curDate)
}

func (u *User) ClearName() {
	u.name = ""
}
