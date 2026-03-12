package main

import "fmt"

func main() {
	// var a int = 1
	// for a != 0 {
	// 	fmt.Print("Enter your choice")

	// 	fmt.Scan(&a)
	// 	if a == 1 {
	// 		fmt.Println("You have entered 1")
	// 	} else if a == 2 {
	// 		fmt.Println("You have entered 2")
	// 	} else {
	// 		fmt.Println("You have entered something else")
	// 	}

	// }

	var str string
	fmt.Scanln(&str)
	fmt.Print("You have entered ", str)
}
