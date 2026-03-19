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

	/*in main function
	 Anoonymous function
	functions without name
	*/

	arr = *makeDoubleByPassingFunc(&arr, func(num int) int {
		return num * 2
	})
	fmt.Println(arr)
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

//recursion function
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
