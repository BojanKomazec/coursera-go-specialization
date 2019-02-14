package main

import "fmt"

func main() {

	var inputValue float32

	fmt.Println("Number of apples")

	num, error := fmt.Scan(&inputValue)

	if 0 == num || nil != error {
		fmt.Println("invalid argument passed")
	} else {
		fmt.Printf("%d", int(inputValue))
	}
}
