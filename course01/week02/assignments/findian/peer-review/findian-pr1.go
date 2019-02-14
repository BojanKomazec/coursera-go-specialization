package main

import "strings"
import "fmt"

func findIan(haystack string) bool {
	sanitized := strings.TrimSpace(strings.ToLower(haystack))
	prefix := strings.HasPrefix(sanitized, "i")
	suffix := strings.HasSuffix(sanitized, "n")
	contains := strings.Contains(sanitized, "a")

	if prefix && suffix && contains {
		return true
	}

	return false
}

func main() {
	var inputValue string

	fmt.Println("Enter your name")

	num, error := fmt.Scan(&inputValue)

	if 0 == num || nil != error {
		fmt.Println("invalid argument passed")
	} else {
		if findIan(inputValue) {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")

		}
	}
}