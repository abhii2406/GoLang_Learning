package main

import (
	"abc/xyz/Test"
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func main() {
	// fmt.Println("Hello, World!")
	var abc string
	var def int
	def, abc = Test.Test("abc", 10)
	fmt.Println(abc, "  ", def)
	fmt.Println(randomdata.PhoneNumber())
	fmt.Println(randomdata.Email())

}
