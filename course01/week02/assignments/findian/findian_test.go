// to run tests use:
//    go test -v

package main

import (
	"testing"
)

func TestIsFound(t *testing.T) {
	var testCases = []struct {
		input          string
		expectedOutput bool
	}{
		{"ian", true},
		{"Ian", true},
		{"iuiygaygn", true},
		{"I d skd a efju N", true},
		{"I A N", true},
		{"ian\n", true},
		{"Ian\n", true},
		{"iuiygaygn\n", true},
		{"I d skd a efju N\n", true},
		{"I A N\n", true},
		{"ian\r\n", true},
		{"Ian\r\n", true},
		{"iuiygaygn\r\n", true},
		{"I d skd a efju N\r\n", true},
		{"I A N\r\n", true},
		{"ihhhhhn", false},
		{"ina", false},
		{"xian", false},
	}

	for _, testCase := range testCases {
		if actualOutput := IsFound(testCase.input); actualOutput != testCase.expectedOutput {
			t.Errorf(
				"Test failed. For %s input, %t was expected but got %t.",
				testCase.input, testCase.expectedOutput, actualOutput,
			)
		}
	}
}
