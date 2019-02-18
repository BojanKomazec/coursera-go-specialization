package main

import (
	"fmt"
	"math"
)

func makeDistOrigin(oX, oY float64) func(float64, float64) float64 {
	fn := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(oX, 2) + math.Pow(oY, 2))
	}
	return fn
}

func testmakeDistOrigin() {
	Dist1 := makeDistOrigin(0, 0)
	Dist2 := makeDistOrigin(2, 2)
	fmt.Println(Dist1(2, 2))
	fmt.Println(Dist2(2, 2))
}

// https://gobyexample.com/closures
func fA() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func testfA() {
	fB := fA()
	fmt.Print(fB())
	fmt.Print(fB())
}

func main() {
	// testmakeDistOrigin()
	// testfA()

	// test defer
	i := 1
	fmt.Print(i)
	i++
	defer fmt.Print(i + 1)
	fmt.Print(i)
	// ~ test
}
