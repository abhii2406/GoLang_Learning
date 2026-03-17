// package main

// import "fmt"

// func main() {
// 	arr := [4]int{1, 4, 2, 3}
// 	fmt.Println(arr)

// 	var names []string
// 	names = append(names, "Abhi")
// 	names = append(names, "Abhi")
// 	// fmt.Println(len(names))

// 	var arr2 [4]int
// 	arr2[0] = 1
// 	arr2[1] = 4
// 	arr2[2] = 2
// 	arr2[3] = 3
// 	// fmt.Println(arr2)

// 	var arr3 [4]int = [4]int{1, 4, 2, 3}
// 	// fmt.Println(arr3)

// 	s := arr3[1:3]
// 	s[0] = 99

// 	// fmt.Println(s)
// 	// fmt.Println(arr3)

// 	// s = arr3[:2]
// 	// fmt.Println(s)

// 	// s = arr3[1:]
// 	// fmt.Println(s)

// 	fmt.Println(len(s))
// 	fmt.Println(cap(arr3))
// }

package main

import "fmt"

func main() {
	arr := [4]int{1, 4, 2, 3}
	fmt.Println(arr)

	s := arr[1:2]

	s1 := s[0:1]

	fmt.Println(s1)
	fmt.Println(len(s1), " ", cap(s1))

}
