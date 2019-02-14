package main

import "fmt"

func main() {
	var floatNumber float64
	fmt.Printf("Enter a floating point number: ")
	fmt.Scan(&floatNumber)
	fmt.Printf("Truncated number: %d\n", int64(floatNumber))
}
