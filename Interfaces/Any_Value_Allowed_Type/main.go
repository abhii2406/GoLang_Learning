package main

import (
	"fmt"
)

func main() {
	AnyTypeValue(1)
	AnyTypeValue("Abhi")
	AnyTypeValue(1.56)
}

func AnyTypeValue(value interface{}) {
	fmt.Println(value)
}
