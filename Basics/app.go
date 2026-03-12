package main

import "fmt"

func main() {
	var revenue int
	var expenses int
	var taxRate float64
	fmt.Print("Revenue")
	fmt.Scan(&revenue)
	fmt.Print("Expenses")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate")
	fmt.Scan(&taxRate)
	fmt.Print("Ebt =", revenue-expenses)
}
