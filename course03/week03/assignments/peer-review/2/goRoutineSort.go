/*
Write a program to sort an array of integers.
The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.
The program should prompt the user to input a series of integers.
Each goroutine which sorts Â¼ of the array should print the subarray that it will sort.
When sorting is complete, the main goroutine should print the entire sorted list.
*/

package main

import (
	"fmt"
	"sort"
	"sync"
)

func input(x []int, err error) []int {
	if err != nil {
		return x
	}
	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		x = append(x, d)
	}
	return input(x, err)
}

func main() {

	var wg sync.WaitGroup
	sorted := []int{}

	fmt.Println("Input a space delimited sequence of integers, press space enter when finished:")
	fmt.Print("> ")
	toSort := input([]int{}, nil)

	// split toSort in 4 slices

	chunk := len(toSort) / 3
	s1 := toSort[0:chunk]
	s2 := toSort[chunk : chunk*2]
	s3 := toSort[chunk*2 : chunk*3]
	s4 := toSort[chunk*3 : len(toSort)]

	wg.Add(4)
	go func() {
		fmt.Println(s1)
		sort.Sort(sort.IntSlice(s1))
		sorted = append(sorted, s1...)
		wg.Done()
	}()

	go func() {
		fmt.Println(s2)
		sort.Sort(sort.IntSlice(s2))
		sorted = append(sorted, s2...)
		wg.Done()
	}()

	go func() {
		fmt.Println(s3)
		sort.Sort(sort.IntSlice(s3))
		sorted = append(sorted, s3...)
		wg.Done()
	}()

	go func() {
		fmt.Println(s4)
		sort.Sort(sort.IntSlice(s4))
		sorted = append(sorted, s4...)
		wg.Done()
	}()

	wg.Wait()
	sort.Sort(sort.IntSlice(sorted))
	fmt.Println(sorted)

}
