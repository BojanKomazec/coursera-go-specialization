package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	storage := make([]int, 3)
	var input string
	counter := 0
	for {
		fmt.Printf("Enter integer: ")
		fmt.Scan(&input)

		if input == "X" {
			break
		}

		if value, err := strconv.Atoi(input); err == nil {
			if counter < 3 {
				storage[counter] = value
			} else {
				storage = append(storage, value)
			}
			sorted := make([]int, len(storage))
			copy(sorted, storage)
			sort.Ints(sorted)
			fmt.Println(sorted)
			counter++
		}
	}
}
