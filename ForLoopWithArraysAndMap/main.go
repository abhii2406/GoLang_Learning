package main

import "fmt"

type intMap map[int]string

func main() {
	names := []string{}
	names = append(names, "Abhi")
	names = append(names, "Ab")

	for i, val := range names {
		fmt.Println(i, " ", val)
	}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(i, " ", names[i])
	// }

	// for _, val := range names {
	// 	fmt.Println(val)
	// }

	m := make(intMap, 2)

	m[0] = "Abhi"
	m[2] = "Ab"

	for key, val := range m {
		fmt.Println(key, "Map", val)
	}
	fmt.Println(len(m))

}
