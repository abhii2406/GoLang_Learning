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
	fun1(abc)

}

func fun1(abc user) {
	fmt.Print(abc.name, " ", abc.age, " ", abc.curDate)
}
