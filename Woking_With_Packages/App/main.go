package main

import (
	"abc/xyz/Test"
	"fmt"
)

func main() {
	// fmt.Println("Hello, World!")
	var abc string
	var def int
	def, abc = Test.Test("abc", 10)
	fmt.Print(abc, "  ", def)

}
