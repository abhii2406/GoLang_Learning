package main

import "fmt"

func main() {
	m := map[string]int{
		"A": 65,
		"B": 66,
	}

	fmt.Println(m)
	fmt.Println(m["A"])

	m["C"] = 67
	fmt.Println(m)
	m["A"] = 64
	fmt.Println(m)
	delete(m, "A")
	fmt.Println(m)
}
