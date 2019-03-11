package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

/*
Write a program to sort an array of integers. The program should partition the array into 4 parts,
each of which is sorted by a different goroutine.
Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers.
Each goroutine which sorts Â¼ of the array should print the subarray that it will sort.
 When sorting is complete, the main goroutine should print the entire sorted list.
*/

func main() {
	inputData := getInputData()
	inputSlice := convertInputDataToIntSlice(inputData)

	if isValidInput(inputSlice) {
		slices := convertInputArrayIntoSubArray(inputSlice)
		fmt.Println("Input Array converted into 4 subarrays:", slices, "\n")

		// use case 1
		fmt.Println("UseCase 1: Sorting 4 sub array by a different goroutine")
		var wg sync.WaitGroup
		fmt.Println("Unsorted First Part: ", slices[0])
		wg.Add(1)
		go BubbleSort(slices[0], &wg)
		wg.Wait()
		fmt.Println("Sorted First Part:   ", slices[0])

		fmt.Println("Unsorted 2nd Part:   ", slices[1])
		wg.Add(1)
		go BubbleSort(slices[1], &wg)
		wg.Wait()
		fmt.Println("Sorted 2nd Part:     ", slices[1])

		fmt.Println("Unsorted 3rd Part:   ", slices[2])
		wg.Add(1)
		go BubbleSort(slices[2], &wg)
		wg.Wait()
		fmt.Println("Sorted 3rd Part:     ", slices[2])

		fmt.Println("Unsorted 4th Part:   ", slices[3])
		wg.Add(1)
		go BubbleSort(slices[3], &wg)
		wg.Wait()
		fmt.Println("Sorted 4th Part:     ", slices[3])

		// use case 2
		fmt.Println("\nUseCase 2: Merge the 4 sorted subarrays into one large sorted Array")
		mergeSubarrays := mergeSubArrays(slices)
		fmt.Println("Consolidated sorted subArray: ", mergeSubarrays)
		wg.Add(1)
		go BubbleSort(mergeSubarrays, &wg)
		wg.Wait()
		fmt.Println("Sorted Consolidated Array:    ", mergeSubarrays)

	}

}

func getInputData() string {
	fmt.Println("Enter  numbers (comma seperated) to be sort. For example: 1,2,3,4")
	reader := bufio.NewReader(os.Stdin)
	inputData, _ := reader.ReadString('\n')
	inputData = strings.TrimSpace(inputData)
	return inputData
}

func isValidInput(inputData []int) bool {
	if len(inputData) < 4 {
		fmt.Println("Minimum 4 interger are required")
		return false
	}
	return true
}

func convertInputDataToIntSlice(inputData string) []int {
	var intSlice []int
	input := strings.Split(inputData, ",")
	for _, value := range input {
		number, _ := strconv.Atoi(value)
		intSlice = append(intSlice, number)
	}
	return intSlice
}

/*
BubbleSort function sort the array
*/
func BubbleSort(input []int, wg *sync.WaitGroup) {
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-1-i; j++ {
			if input[j] > input[j+1] {
				Swap(input, j)
			}
		}
	}
	wg.Done()
}

/*
Swap swaps two number
*/
func Swap(input []int, index int) {
	noAtIndexPosition := input[index]
	input[index] = input[index+1]
	input[index+1] = noAtIndexPosition
}

func mergeSubArrays(slices [4][]int) []int {
	var tmp []int
	for _, s := range slices {
		tmp = append(tmp, s...)
	}
	return tmp
}

func convertInputArrayIntoSubArray(inputSlice []int) [4][]int {
	i := 0
	var slices [4][]int
	for _, v := range inputSlice {
		slices[i] = append(slices[i], v)
		i++
		if i == 4 {
			i = 0
		}
	}
	return slices
}
