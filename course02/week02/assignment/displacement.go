package main

import (
	"fmt"
)

func inputFloat64(valueName string) float64 {
	var number float64

	for {
		fmt.Print("Please enter a float64 value for ", valueName, ": ")
		var err error
		if _, err = fmt.Scanln(&number); err == nil {
			break
		} else {
			fmt.Println("\nError:", err)
		}
	}

	return number
}

// GenDisplaceFn generates a function which calculates displacement for the given time respecting acceleration, initial
// velocity and and intial displacement set from this function's arguments
func GenDisplaceFn(acceleration float64, initialVelocity float64, initialDisplacement float64) func(time float64) float64 {
	return func(time float64) float64 {
		return 0.5*acceleration*time*time + initialVelocity*time + initialDisplacement
	}
}

func main() {
	acceleration := inputFloat64("acceleration")
	initialVelocity := inputFloat64("inial velocity")
	initialDisplacement := inputFloat64("initial displacement")
	computeDisplacement := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	time := inputFloat64("time")
	displacement := computeDisplacement(time)
	fmt.Println("result (displacement) =", displacement)
}
