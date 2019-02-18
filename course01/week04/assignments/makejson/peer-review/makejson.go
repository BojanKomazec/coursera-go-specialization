package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a name:")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error while reading name: %v\n", err)
		return
	}
	name = strings.TrimRight(name, "\r\n")

	fmt.Println("Enter an address:")
	address, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error while reading address: %v\n", err)
		return
	}
	address = strings.TrimRight(address, "\r\n")

	obj := map[string]string{}
	obj["name"] = name
	obj["address"] = address

	jsonData, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Println(string(jsonData))
}
