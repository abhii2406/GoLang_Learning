package main

import (
	"fmt"
)

func main() {
	AnyTypeValue(1)
	AnyTypeValue("Abhi")
	AnyTypeValue(1.56)
}

// func AnyTypeValue(value interface{}) {
// 	// fmt.Println(value)
// 	switch value.(type) {
// 	case int:
// 		fmt.Println("Int : ", value)
// 	case float64:
// 		fmt.Println("Float64 : ", value)
// 	case string:
// 		fmt.Println("String : ", value)
// 	}
// }

func AnyTypeValue(value interface{}) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer : ", intVal)
		return
	}
	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float : ", floatVal)
		return
	}
	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String : ", stringVal)
		return
	}
}
