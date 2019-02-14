package main

import (
	"fmt"
)

func arrayDemo() {
	x := [3]int{1, 2, 3}
	for i, v := range x {
		fmt.Printf("index = %d; element value = %d", i, v)
	}
}

func arrayIterationDemo() {
	x := []int{4, 8, 5}
	y := -1
	for _, elt := range x {
		if elt > y {
			y = elt
		}
	}
	fmt.Print(y)
}

func printSlice(slice []int) {
	for _, value := range slice {
		fmt.Printf("%d, ", value)
	}
}

func sliceDemo() {
	a2 := [3]string{"a", "b", "c"}
	fmt.Println("a2 = ", a2)
	slice21 := a2[0:1] // { "a" }
	fmt.Println("slice21 = ", slice21)
	fmt.Println(len(slice21), cap(slice21)) // 1, 3

	slice3 := make([]int, 10)
	fmt.Println("slice3 length = ", len(slice3)) // 10
}

func sliceDemo2() {
	x := [...]int{4, 8, 5}
	y := x[0:2]
	z := x[1:3]
	y[0] = 1
	z[1] = 3
	fmt.Print(x)
}

func sliceLengthCapacityDemo() {
	x := [...]int{1, 2, 3, 4, 5}
	y := x[0:2]
	z := x[1:4]
	fmt.Print(len(y), cap(y), len(z), cap(z))
}

func sliceLengthCapacityDemo2() {
	s := make([]int, 0, 3)
	s = append(s, 100)
	fmt.Println(len(s), cap(s))
}

func appendDemo() {
	s1 := []int{0, 1, 2, 3, 4}
	fmt.Println("s1 =", s1)

	s2 := s1[0:3]
	fmt.Println("s2 =", s2) // [0 1 2]

	s2 = append(s2, 9)
	fmt.Println("s1 =", s1) // [0 1 2 9 4]
	fmt.Println("s2 =", s2) // [0 1 2 9]
}

func circularShiftRight(s []int) []int {
	length := len(s)
	last := s[length-1]
	for index := length - 1; index > 0; index-- {
		s[index] = s[index-1]
	}
	s[0] = last
	return s
}

func circularShiftRightDemo() {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 = circularShiftRight(s1)
	fmt.Println("s1 =", s1)

	s1 = []int{0, 1, 2, 3, 4, 5, 6, 7}
	s2 := s1[3:len(s1)] // [3 4 5 6 7]
	fmt.Println("s2 =", s2)

	s2 = circularShiftRight(s2)
	fmt.Println("s1 =", s1) // [0 1 2 7 3 4 5 6]
	fmt.Println("s2 =", s2) // [7 3 4 5 6]
}

func mapIterationDemo() {
	x := map[string]int{"ian": 1, "harris": 2}

	for i, j := range x {
		if i == "harris" {
			fmt.Print(i, j)
		}
	}
}

type P struct {
	x string
	y int
}

func structDemo() {
	b := P{"x", -1}
	a := [...]P{
		P{"a", 10},
		P{"b", 2},
		P{"c", 3},
	}

	for _, z := range a {
		if z.y > b.y {
			b = z
		}
	}

	fmt.Println(b.x)
}

func main() {
	// sliceDemo()
	// appendDemo()
	// circularShiftRightDemo()
	// arrayIterationDemo()
	// sliceDemo2()
	// sliceLengthCapacityDemo()
	// mapIterationDemo()
	// structDemo()
	sliceLengthCapacityDemo2()
}
