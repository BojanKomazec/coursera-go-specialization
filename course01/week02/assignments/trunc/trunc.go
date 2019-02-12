package main

import "fmt"

func main() {
	fmt.Printf("Please enter a floating point number: ")
	var f float64
	_, err := fmt.Scan(&f)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		n := int64(f)
		fmt.Printf("Truncated value is: %d", n)
	}
}
