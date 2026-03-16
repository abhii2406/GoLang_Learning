package main

import "fmt"

func main() {
	fmt.Println(add(10, 20))
	fmt.Println(add(10.5, 20))
	fmt.Println(add("A", "B"))
}

// func add[T any](a,b T)T{
// 	return  a+b    //any Creates a prolem  -->invalid operation: operator + not defined on a (variable of type T constrained by any)compilerUndefinedOp
// 				   // returns (T)
// }

func add[T int | float64 | string](a, b T) T {
	return a + b
}
