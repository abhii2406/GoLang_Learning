package main

import (
	"learning/module/Structs/Demo4/user"
	"time"
)

func main() {

	var abc *user.User
	abc = user.NewUser("Abhi", 23, time.Now())
	abc.PrintAll()
	abc.ClearName()
	abc.PrintAll()

}
