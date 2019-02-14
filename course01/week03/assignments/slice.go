package main

import (
	"fmt"
	"strconv"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func retrieveInput() (string, error) {
	fmt.Println("Please enter an integer or X to EXIT: ")
	var input string
	if _, err := fmt.Scan(&input); err != nil {
		return "", err
	} else {
		return input, nil
	}
}

// CircularShiftRight circulary shifts all elements in slice to the right by one place.
func CircularShiftRight(s []int) []int {
	// fmt.Println("circularShiftRight(): input slice: ", s)
	length := len(s)
	last := s[length-1]
	for index := length - 1; index > 0; index-- {
		s[index] = s[index-1]
	}
	s[0] = last
	// fmt.Println("circularShiftRight(): output slice: ", s)
	return s
}

// InsertToSorted inserts new element to the right place in a slice of sorted elements and returns sorted slice.
func InsertToSorted(slice []int, newElement int) []int {
	// fmt.Println("InsertToSorted(): input slice: ", slice, "input value:", newElement)

	if len(slice) == 0 {
		slice = append(slice, newElement)
	} else {
		slice = append(slice, newElement)
		// fmt.Println("InsertToSorted(): slice (after appending new): ", slice)

		var indexOfFirstLarger = -1
		for index, element := range slice {
			if element > newElement {
				indexOfFirstLarger = index
				break
			}
		}

		if indexOfFirstLarger > -1 {
			CircularShiftRight(slice[indexOfFirstLarger:len(slice)])
		}
	}
	// fmt.Println("InsertToSorted(): output slice: ", slice)
	return slice
}

func main() {
	slice := make([]int, 3) // create an empty integer slice of size (length) 3
	var input string

	for input != "X" {
		if input, err := retrieveInput(); err != nil {
			fmt.Println("Input error:", err, ".Please try again.")
			continue
		} else {
			fmt.Println("Input string:", input)
			if input == "X" {
				break
			}

			if number, err := strconv.Atoi(input); err != nil {
				fmt.Println("Input string is not an integer number")
				continue
			} else {
				fmt.Println("Input string is an integer number: ", number)
				slice = InsertToSorted(slice, number)
				fmt.Println("Slice in the (increasing) order:", slice)
			}
		}
	}
}
