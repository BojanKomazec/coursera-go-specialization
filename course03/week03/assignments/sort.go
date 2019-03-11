package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// ReadStringLine reads a set of characters before ENTER is hit from a standard input (terminal)
func ReadStringLine() (string, error) {
	stdinReader := bufio.NewReader(os.Stdin)

	if inputString, err := stdinReader.ReadString('\n'); err != nil {
		fmt.Println("Invalid input. Error: ", err)
		return "", err
	} else {
		// if exists, remove trailing Line Feed (added on Windows and on Unix/OSX)
		if strings.HasSuffix(inputString, "\n") {
			inputString = strings.TrimSuffix(inputString, "\n")
		}

		// if exist, remove trailing Carriage Return (added on Windows)
		if strings.HasSuffix(inputString, "\r") {
			inputString = strings.TrimSuffix(inputString, "\r")
		}

		inputString = strings.TrimSpace(inputString)

		// fmt.Println("inputString = ", inputString)
		return inputString, nil
	}
}

// ToIntegers takes a string, splits it in chunks separated by SPACE, converts each chunk into an integer and returns an
// array of integers.
func ToIntegers(str string) ([]int, error) {
	parts := strings.Split(str, " ")
	slice := make([]int, 0, len(parts))
	for _, v := range parts {
		if len(v) == 0 {
			continue
		}
		if n, err := strconv.Atoi(v); err != nil {
			fmt.Println("Error:", err)
			return nil, err
		} else {
			slice = append(slice, n)
		}
	}
	return slice, nil
}

// ReadIntegersLine reads characters before ENTER is hit from a standard input and returns an array of integers.
func ReadIntegersLine() ([]int, error) {
	if inputString, err := ReadStringLine(); err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	} else {
		// fmt.Println("inputString = ", inputString)
		return ToIntegers(inputString)
	}
}

func printElements(arr []int) {
	// fmt.Println("Number of elements:", len(arr))
	if len(arr) > 0 {
		// fmt.Println("Elements: ")
		for _, v := range arr {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	} else {
		fmt.Println("No elements found!")
	}
}

// Swap function swaps i-th and i+1-th elements in the input sequence.
func SwapAdjacent(sequence []int, index int) {
	if index >= len(sequence)-1 {
		return
	}
	temp := sequence[index]
	sequence[index] = sequence[index+1]
	sequence[index+1] = temp
}

// Swap function swaps i-th and i+1-th elements in the input sequence.
func Swap(sequence []int, index int, index2 int) {
	if index > len(sequence)-1 {
		return
	}

	if index2 > len(sequence)-1 {
		return
	}

	temp := sequence[index]
	sequence[index] = sequence[index2]
	sequence[index2] = temp
}

// BubbleSort uses Bubble Sort algorithm to sort input sequence of numbers.
// This function only has to modify elements in the slice but not to change slice's size/length => we can pass []int
func BubbleSort(sequence []int, wg *sync.WaitGroup) {
	fmt.Println("Sequence to be sorted:", sequence)
	lenght := len(sequence)

	for i := 0; i < lenght; i++ {
		for j := 0; j < lenght-1-i; j++ {
			if sequence[j] > sequence[j+1] {
				SwapAdjacent(sequence, j)
			}
		}
	}

	if wg != nil {
		wg.Done()
	}
}

// 5 34 3 6 2 1 88 34 21 1 23 9 44 25 5 1
// 5 34 3 6  => 3 5 6 34
// 2 1 88 34 => 1 2 34 88
// 21 1 23 9 => 1 9 21 23
// 44 25 5 1 => 1 5 25 44

// 3  5   6  34
// 1  2  34  88
// 3 > 1 => s2[0] goes to result slice, j++ (j == 1); x == 0

// 3  5   6  34
// x  2  34  88
// 3 > 2 => s2[1] goes to result slice, j++ (j == 2); x == 0

// 3  5   6  34
// x  x  34  88
// 3 < 34 => s1[0] goes to result slice, i++ (i == 1); j == 2

// x  5   6  34
// x  x  34  88

// etc...

func MergeTwoSorted(slice1 []int, slice2 []int, slice12 []int) {
	fmt.Println()
	// fmt.Println("MergeTwoSorted():", len(slice1), len(slice2), len(slice12))
	// fmt.Println("MergeTwoSorted():", slice1, slice2, slice12)
	i, j, k := 0, 0, 0
	for {
		// fmt.Println("slice1[", i, "]=", slice1[i])
		// fmt.Println("slice2[", j, "] =", slice2[j])
		if slice1[i] < slice2[j] {
			slice12[k] = slice1[i]
			k++
			i++
		} else if slice1[i] == slice2[j] {
			slice12[k] = slice1[i]
			k++
			i++

			slice12[k] = slice2[j]
			k++
			j++
		} else if slice1[i] > slice2[j] {
			slice12[k] = slice2[j]
			k++
			j++
		}

		if i == len(slice1) {
			for j < len(slice2) {
				slice12[k] = slice2[j]
				k++
				j++
			}
		}

		if j == len(slice2) {
			for i < len(slice1) {
				slice12[k] = slice1[i]
				k++
				i++
			}
		}

		if k == len(slice12) {
			break
		}
	}

	// fmt.Println("MergeTwoSorted():", slice12)
}

func main() {
	fmt.Println("Please enter a sequence of numbers to be sorted and then press ENTER:")
	//
	// Read user input - a sequence of integers
	//
	if intSlice, err := ReadIntegersLine(); err != nil {
		fmt.Println("Error:", err)
	} else {
		// printElements(intSlice)

		elementsTotalCount := len(intSlice)

		if elementsTotalCount < 4 {
			BubbleSort(intSlice, nil)
			fmt.Println("All numbers sorted:", intSlice)
		} else {
			//
			// Create 4 partitions
			//
			elementsChunkCount := elementsTotalCount / 4
			slice1 := intSlice[0:elementsChunkCount]
			slice2 := intSlice[elementsChunkCount : 2*elementsChunkCount]
			slice3 := intSlice[2*elementsChunkCount : 3*elementsChunkCount]
			slice4 := intSlice[3*elementsChunkCount : elementsTotalCount]

			//
			// Sort each partition in parallel
			//
			// If executing the applicaton multiple times with the same sequence of numbers (with more than 4 numbers)
			// observe how each goroutine would take the same partition and on each running they would be run in
			// different order (their order is undeterminstic and depends on Go runtime scheduler) and their sequences
			// would be print in different order.
			//
			var wg sync.WaitGroup
			wg.Add(4)
			go BubbleSort(slice1, &wg)
			go BubbleSort(slice2, &wg)
			go BubbleSort(slice3, &wg)
			go BubbleSort(slice4, &wg)
			wg.Wait()

			// fmt.Println("Partitions after sorting:")
			// printElements(slice1)
			// printElements(slice2)
			// printElements(slice3)
			// printElements(slice4)
			// fmt.Println("Slice after sorting partitions:")
			// printElements(intSlice)

			slice12 := make([]int, len(slice1)+len(slice2))
			MergeTwoSorted(slice1, slice2, slice12)

			slice34 := make([]int, len(slice3)+len(slice4))
			MergeTwoSorted(slice3, slice4, slice34)

			slice1234 := make([]int, len(slice12)+len(slice34))
			MergeTwoSorted(slice12, slice34, slice1234)

			fmt.Println("All numbers sorted:", slice1234)
		}
	}
}
