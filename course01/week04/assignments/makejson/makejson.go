package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Person struct
// type Person struct {
// 	name    string
// 	address string
// }

// Person is a map with two keys "name" and "address"
type Person map[string]string

func trimTrailingCrLf(s string) string {
	trimmed := s
	if strings.HasSuffix(trimmed, "\n") {
		trimmed = strings.TrimSuffix(trimmed, "\n")
	}
	if strings.HasSuffix(trimmed, "\r") {
		trimmed = strings.TrimSuffix(trimmed, "\r")
	}

	return trimmed
}

// Using buffer reader allows both name and address to contain SPACE characters.
func inputPersonDetails() Person {
	person := make(Person)
	stdinReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Please enter the name:")
		if name, err := stdinReader.ReadString('\n'); err != nil {
			fmt.Println("\nError: ", err)
		} else {
			// person.name = trimTrailingCrLf(name)
			person["name"] = trimTrailingCrLf(name)
			break
		}
	}

	for {
		fmt.Print("Please enter the address:")
		if address, err := stdinReader.ReadString('\n'); err != nil {
			fmt.Println("\nError: ", err)
		} else {
			// person.address = trimTrailingCrLf(address)
			person["address"] = trimTrailingCrLf(address)
			break
		}
	}

	return person
}

func main() {
	person := inputPersonDetails()
	// fmt.Println("\nName:", person["name"], "\nAddress:", person["address"])

	// For pretty-printing use MarshalIndent:
	// if data, err := json.MarshalIndent(person, "", "   "); err != nil {

	if data, err := json.Marshal(person); err != nil {
		fmt.Println("\nError in serializing to JSON: ", err)
	} else {
		fmt.Println("\nJSON:", string(data))
	}
}
