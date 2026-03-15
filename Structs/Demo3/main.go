package main

import (
	"fmt"
	"time"
)

type user struct {
	name    string
	age     int
	curDate time.Time
}

// func newUser(name string, age int, curDate time.Time) user {
// 	return user{                   // creating unnecessary object of user type , it avoid during cop;ex application
// 		name:    name,
// 		age:     age,
// 		curDate: curDate,
// 	}
// }

func newUser(name string, age int, curDate time.Time) *user {
	return &user{
		name:    name,
		age:     age,
		curDate: curDate,
	}
}

func main() {

	var abc *user
	abc = newUser("Abhi", 23, time.Now())
	abc.printAll()
	abc.clearName()
	abc.printAll()

}

func (u user) printAll() {
	fmt.Println(u.name, " ", u.age, " ", u.curDate)
}

func (u *user) clearName() {
	u.name = ""
}
