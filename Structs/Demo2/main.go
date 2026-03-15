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

func main() {

	var abc user
	abc = user{
		name:    "Abhi",
		age:     23,
		curDate: time.Now(),
	}
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
