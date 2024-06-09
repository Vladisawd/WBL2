package main

import (
	"testing"
)

func Test_Sort(t *testing.T) {
	testTable := []struct {
		input  []string
		output []string
		flags  flags
	}{
		{
			input:  []string{"2000", "1000", "10000", "33234", "53423"},
			output: []string{"1000", "10000", "2000", "33234", "53423"},
			flags:  flags{column: -1, byNumber: true},
		},
		{
			input:  []string{"a", "a", "b", "c", "d"},
			output: []string{"a", "b", "c", "d"},
			flags:  flags{column: -1, noRepetitive: true},
		},
		{
			input:  []string{"1, 3, 4", "2, 2, 5", "3, 1, 6"},
			output: []string{"3, 1, 6", "2, 2, 5", "1, 3, 4"},
			flags:  flags{column: 1},
		},
		{
			input:  []string{"1", "2", "3", "4", "5"},
			output: []string{"5", "4", "3", "2", "1"},
			flags:  flags{column: -1, inReverseOrder: true},
		},
	}

	for _, testCase := range testTable {
		result := flagsSort(testCase.input, &testCase.flags)
		for i := 0; i < len(result); i++ {
			if result[i] != testCase.output[i] {
				t.Errorf("Incorrect result. Expect: %v, got %v", testCase.output[i], result[i])
			}
		}
	}
}
