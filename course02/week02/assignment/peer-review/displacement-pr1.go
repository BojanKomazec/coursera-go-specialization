package main

import (
	"fmt"
)

func main() {
	a := readInput("Acceleration:")
	v0 := readInput("Initial Velocity:")
	s0 := readInput("Initial Displacement:")

	fn := GenDisplaceFn(a, v0, s0)

	t := readInput("Time:")

	fmt.Println(fn(t))
}

func readInput(msg string) float64 {
	fmt.Print(msg)
	var result float64
	fmt.Scanln(&result)
	return result
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}
