package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5}

	// arr = *makeDouble(&arr)
	// fmt.Println(arr)
	arr = *makeDoubleByPassingFunc(&arr, double)
	fmt.Println(arr)

	getD := getDouble()
	fmt.Println(getD(2))

}

func makeDouble(arr *[]int) *[]int {
	res := []int{}
	for _, val := range *arr {
		res = append(res, double(val))
	}

	return &res
}

// we can pass functions as value like first class citizen

func makeDoubleByPassingFunc(arr *[]int, double func(int) int) *[]int {
	res := []int{}
	for _, val := range *arr {
		res = append(res, double(val))
	}
	return &res
}

func double(a int) int {
	return a * 2
}

// also function can returnibg another function

func getDouble() func(int) int {
	return double
}
