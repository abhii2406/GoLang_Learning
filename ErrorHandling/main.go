package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	x, err := userInput("X")
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	fmt.Println(x)
	y, err := userInput("Y")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(y)
	z, err := userInput("Z")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(z)

	storeResultInFile(x, y, z)
}

func userInput(str string) (int, error) {
	var x int
	fmt.Println("Enter Value for ", str, " ")
	fmt.Scan(&x)
	if x <= 0 {
		return 0, errors.New("Value Must me Greater Than 0")
	}
	return x, nil
}

func storeResultInFile(x, y, z int) {
	result := fmt.Sprintf("Value Of X : %d, \n Value Of Y : %d, \n Value Of Z : %d", x, y, z)
	os.WriteFile("result.txt", []byte(result), 0644)
}
