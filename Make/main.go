package main

import "fmt"

func main() {
	// arrSlice := make([]string, 2) // it creates dynamic array with first two null i.e empty strings
	// arrSlice[0] = "Abhi"
	// arrSlice = append(arrSlice, "Abhi")
	// fmt.Println(arrSlice)

	arrSlice := make([]string, 2, 5) // 5 denoting the slice can hold 5 value without resizing
	arrSlice = append(arrSlice, "A")
	arrSlice = append(arrSlice, "B")
	arrSlice = append(arrSlice, "H")
	arrSlice = append(arrSlice, "I")
	arrSlice = append(arrSlice, "J")
	fmt.Println(arrSlice)

	/*
		append "A" → ["", "", "A"]
		append "B" → ["", "", "A", "B"]
		append "H" → ["", "", "A", "B", "H"]
		append "I" → capacity exceeded → new array created
		append "J"
	*/

	m := make(map[string]float64, 3)
	m["A"] = 65.5
	m["B"] = 66.5
	m["C"] = 67.5
	m["D"] = 68.5 // here resizing happens when we add 4th element
	fmt.Println(m)
}
