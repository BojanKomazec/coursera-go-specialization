package main

import (
	"fmt"
	"strings"
)

func inputNumber() int {
	var number int

	for {
		fmt.Print("Please enter a number: ")
		var err error
		if _, err = fmt.Scanln(&number); err == nil {
			break
		} else {
			fmt.Println("Error:", err)
		}
	}

	return number
}

func inputYesNo() string {
	var answer string
	for {
		fmt.Print("Would you like to add more numbers?(y/n)")
		var err error
		if _, err = fmt.Scanln(&answer); err == nil {
			answer = strings.ToLower(answer)
			if answer == "y" || answer == "n" {
				break
			}
		} else {
			fmt.Println("Error:", err)
		}
	}
	return answer
}

// [10]int would mean that type of the argument is an array => function gets a copy of it so can't modify it!
// []int means that the type or argument is a slice => function can modify its elements but not size!
// *[]int means that the type or argument is a pointer to slice => function can modify its size and its elements
func inputNumbers(sequence *[]int) {
	const maxNumbers = 10
	for index := 0; index < maxNumbers; index++ {
		*sequence = append(*sequence, inputNumber())
		// fmt.Println("Sequence after adding a number:", *sequence)
		if index == maxNumbers-1 {
			fmt.Println("You've entered maximum (10) number of integers.")
		} else {
			answer := inputYesNo()
			if answer == "n" {
				break
			}
		}
	}
}

// Swap function swaps i-th and i+1-th elements in the input sequence.
func Swap(sequence []int, index int) {
	if index >= len(sequence)-1 {
		return
	}
	temp := sequence[index]
	sequence[index] = sequence[index+1]
	sequence[index+1] = temp
}

// BubbleSort uses Bubble Sort algorithm to sort input sequence of numbers.
// This function only has to modify elements in the slice but not to change slice's size/length => we can pass []int
func BubbleSort(sequence []int) {
	lenght := len(sequence)

	for i := 0; i < lenght; i++ {
		for j := 0; j < lenght-1-i; j++ {
			if sequence[j] > sequence[j+1] {
				Swap(sequence, j)
			}
		}
	}
}

func main() {
	fmt.Println("Welcome to BubbleSort demo! You will be prompted to enter up to 10 numbers which will be sorted in the increasing order")
	sequence := make([]int, 0, 10)
	inputNumbers(&sequence)
	fmt.Println("Original sequence:", sequence)
	BubbleSort(sequence)
	fmt.Println("Sorted sequence:", sequence)
}
