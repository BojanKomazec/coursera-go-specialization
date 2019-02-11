package main

import "fmt"

func foo() *int {
	i := 1
	return &i
}

func main() {
	var y *int
	y = foo()
	fmt.Printf("*y = %d", *y) // expected: *y = 1
}
