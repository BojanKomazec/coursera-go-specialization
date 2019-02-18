package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	slice := ReadSlice()

	BubbleSort(slice)

	fmt.Println(slice)
}

func BubbleSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		for j := 0; j < len(slice)-i; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

func Swap(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}

func ReadSlice() []int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter a size of array from 0 to 10: ")
	scanner.Scan()
	size, err := strconv.Atoi(scanner.Text())
	if err != nil || size < 0 || size > 10 {
		log.Fatalln("Size is not correct")
	}
	fmt.Println("Please enter an array to sort (press enter after each element of the array):")
	result := make([]int, size)
	for i := 0; i < size; i++ {
		scanner.Scan()
		item, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln("Passed icorrect integer")
		}
		result[i] = item
	}
	return result
}
