package main

import (
	"fmt"
	"os"
)

func main() {
	var str string = "abhijit"
	writeToFile(str)
	data := readFromFile()
	fmt.Print("Data is ", data)
}

func writeToFile(str string) {
	os.WriteFile("demo.txt", []byte(str), 0644)
}

func readFromFile() string {
	data, _ := os.ReadFile("demo.txt")
	return string(data)

}
