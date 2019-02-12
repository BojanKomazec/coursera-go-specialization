// https://golang.org/doc/faq#stack_or_heap
// "if the compiler cannot prove that the variable is not referenced after the function returns, then the compiler must
// allocate the variable on the garbage-collected heap to avoid dangling pointer errors."

package main

import (
	"fmt"
	"strconv"
	"strings"
)

var g = 4

func printGlobal() {
	fmt.Printf("g = %d", g)
}

func printGlobalAndLocal() {
	fmt.Printf("g = %d", g)
	var l = 5
	fmt.Printf("l = %d", l)
}

// i will not be deallocated as long as there is some reference to it (in this case it's in main())
//
// http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#stack_heap_vars:
// In Go the compiler decides where the variable will be allocated even if the new() or make() functions are used.
// The compiler picks the location to store the variable based on its size and the result of "escape analysis".
// This also means that it's ok to return references to local variables, which is not ok in other languages like C or
// C++.
func foo() *int {
	i := 1
	return &i
}

func usePointerToFunctionLocalVariable() {
	var y *int
	y = foo()
	fmt.Printf("*y = %d\n", *y) // expected: *y = 1
}

func printDemo() {
	name := "Bojan"

	// concatenation operator
	fmt.Printf("Hello, " + name + "\n")

	// formatting directives (%d, %s...) can't be used with fmt.Printl
	// https://stackoverflow.com/questions/53961617/call-has-possible-formatting-directive
}

func conversionDemo() {
	var x int32 = 1
	var y int16 = 2

	// [go] cannot use y (type int16) as type int32 in assignment
	// x = y
	x = int32(y)
	fmt.Printf("x = %d ", x)

	var a, b int = 3, 4
	avg := float64(a+b) / 2    // conversion always goes to the wider type
	fmt.Println("avg = ", avg) // avg = 3.5
	avg = float64(a+b) / 2.0
	fmt.Println("avg = ", avg) // avg = 3.5
	avg = float64(float64(a+b) / 2.0)
	fmt.Println("avg = ", avg) // avg = 3.5
}

// import "strconv" required
func convStringToNumberDemo() {
	i, _ := strconv.Atoi("10")
	y := i * 2
	fmt.Println(y)
}

func constDemo() {
	const x = 1
	const (
		y = 2
		s = "Bojan"
	)
}

// in the current implementation iota values start from 1 but this is not guaranteed and can change in future
func iotaDemo() {
	type Grades int
	const (
		A Grades = iota
		B
		C
		D
		E
		F
	)

	fmt.Printf("C = %d", C) // output: C = 2
}

func forLoopDemo() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d", i)
		if i == 5 {
			break
		}
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Printf("i = %d", i)
	}

	j := 0
	for j < 10 {
		fmt.Printf("j = %d", j)
		j++
	}

	for {
		fmt.Printf("infinite loop")
	}
}

func forLoopDemo2() {
	var xtemp int
	x1 := 0
	x2 := 1
	for x := 0; x < 5; x++ {
		xtemp = x2
		x2 = x2 + x1
		x1 = xtemp
	}
	fmt.Println(x2)
}

// no need to use 'break' like in C; execution breaks automatically
func switchDemo() {
	x := 2
	switch x {
	case 1:
		fmt.Printf("case 1")
	case 2:
		fmt.Printf("case 2")
	default:
		fmt.Printf("default")
	}
}

// cases don't need break statement - it is implicit!
func taglessSwitchDemo() {
	x := 2
	switch {
	case x < -1:
		fmt.Printf("case 1")
	case x > 1:
		fmt.Printf("case 2")
	default:
		fmt.Printf("default")
	}

	fmt.Println("\n2nd example: x == 7:")

	x = 7
	switch {
	case x > 3:
		fmt.Printf("1")
	case x > 5:
		fmt.Printf("2")
	case x == 7:
		fmt.Printf("3")
	default:
		fmt.Printf("4")
	}
}

func inputNumber() {
	var n int
	fmt.Printf("Type in some number: ")
	_, err := fmt.Scan(&n) // although ENTER stops the input, ENTER is not taken so it remains in input buffer!
	if err == nil {
		// fmt.Printf("n = %d", n)
		fmt.Println("n = ", n)
	}
}

func scanStringDemo() {
	var str string
	fmt.Print("scanStringDemo(): Type in some string: ")
	fmt.Scan(&str) // uses SPACE as value delimiter, although ENTER stops the input, ENTER is not taken so it remains in input buffer!
	fmt.Println("Typed string is: ", str)
}

func scanfStringDemo() {
	var str string

	fmt.Print("Type in some string: ")
	fmt.Scanf("%s", &str)
	fmt.Println("Typed string is: ", str)
}

func scanlnStringDemo() {
	var str string

	fmt.Print("scanlnStringDemo(): Type in some string: ")
	fmt.Scanln(&str) // ENTER terminates input
	fmt.Println("Typed string is: ", str)
}

func stringsDemo() {
	s := strings.Replace("ianianian", "ni", "in", 2) // expected: iainainan
	fmt.Println(s)
}

func scanDemo() {
	//inputNumber()
	//scanStringDemo()
	//scanfStringDemo()
	scanlnStringDemo()
}

// func pointersDemo() {
// 	var x int
// 	var y *int
// 	z := 3
// 	y = &z
// 	//x = &y
// }

func main() {
	// conversionDemo()
	// usePointerToFunctionLocalVariable()
	// printDemo()
	// iotaDemo()
	// scanDemo()
	convStringToNumberDemo()
	// stringsDemo()
	// taglessSwitchDemo()
	forLoopDemo2()
}
