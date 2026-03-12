package main

import "fmt"

func main() {
	// fmt.Println("Hello, World!")
	var abc string
	var def int
	def, abc = test("abc", 10)
	fmt.Print(abc, "  ", def)

}

func test(abc string, def int) (int, string) {
	var a int = 10
	var b string = "abc"

	return a, b
}
