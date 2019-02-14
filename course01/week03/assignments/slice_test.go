package main

import (
	"reflect"
	"testing"
)

type TestCaseData struct {
	inputSlice   []int
	inputElement int
	outputSlice  []int
}

func TestInsertToSorted(t *testing.T) {
	var testCases = []TestCaseData{
		{[]int{}, 1, []int{1}},
		{[]int{2}, 1, []int{1, 2}},
	}

	for index, testCase := range testCases {
		actualOutput := InsertToSorted(testCase.inputSlice, testCase.inputElement)
		if !reflect.DeepEqual(actualOutput, testCase.outputSlice) {
			t.Errorf(
				"%d. test failed.",
				index,
			)
		}
	}
}
