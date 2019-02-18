// to run tests use:
//    go test -v

package main

import (
	"testing"
)

// s = 1/2 * a * t^2 + v0 * t + s0
func TestGenDisplaceFn(t *testing.T) {
	var testCases = []struct {
		acceleration         float64
		initialVelocity      float64
		initialDisplacement  float64
		time                 float64
		expectedDisplacement float64
	}{
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 2.5},
		// {0.235, 1.345, 98.07, 5.62, 109.340067}, // %f returns 109.340067; actual value is 109.34006699999999
		{1.435, 6.456245, 18.07, 25.459832, 647.5305981742907}, // 465.0856851234507 + 164.37491305084 + 18.07
	}

	for _, testCase := range testCases {
		computeDisplacement := GenDisplaceFn(testCase.acceleration, testCase.initialVelocity, testCase.initialDisplacement)
		if actualDisplacement := computeDisplacement(testCase.time); actualDisplacement != testCase.expectedDisplacement {
			t.Errorf(
				"Test failed. For %f acceleration, %f initialVelocity, %f initialDisplacement, %f time, %f was expected but got %f.",
				testCase.acceleration, testCase.initialVelocity, testCase.initialDisplacement, testCase.time, testCase.expectedDisplacement, actualDisplacement,
			)
		}
	}
}
