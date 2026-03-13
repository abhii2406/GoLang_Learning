package main

import "fmt"

func main() {
	a := "Abhi"
	// b := "Abhi"
	var p *string
	p = &a
	q := *p
	test(*p, q)
	x := 10
	ptr := &x
	fmt.Println(x)
	test1(ptr)
	fmt.Println(x)
	fmt.Println(&x, "   ", ptr)
}

func test(a string, q string) {
	// fmt.Print(&a, "\n", &q)
}

func test1(ptr *int) {
	*ptr = *ptr + 10
}
