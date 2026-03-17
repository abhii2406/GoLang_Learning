package main

import "fmt"

type Products struct {
	id    string
	title string
	price float64
}

func main() {
	arr := [3]string{"reading", "swimming", "travelling"}

	fmt.Println(arr)

	fmt.Println(arr[0])
	fmt.Println(arr[1:3])

	arrSlice := arr[:2]
	fmt.Println(arrSlice)

	// arrSlice = arrSlice[1:]
	arrSlice = arrSlice[1:3]
	fmt.Println(arrSlice)

	goals := []string{"Learn Go", "Learn Basics"}
	fmt.Println(goals)

	goals[0] = "Learn"
	goals = append(goals, "Learn Go")
	fmt.Println(goals)

	// products := []products{
	// 	products{id: "1", title: "abc", price: 10.55},
	// 	products{id: "2", title: "xyz", price: 20.55},
	// }
	products := []Products{
		{id: "1", title: "abc", price: 10.55},
		{id: "2", title: "xyz", price: 20.55},
	}
	products = append(products, Products{id: "3", title: "pqr", price: 30.55})

	p := Products{
		"4",
		"def",
		40.55,
	}
	products = append(products, p)
	fmt.Println(products)

	array := []int{1, 2, 3}
	array1 := []int{4, 5, 6}

	array = append(array, array1...)
	fmt.Println(array)
}
