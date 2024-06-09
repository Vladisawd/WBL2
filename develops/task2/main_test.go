package main

import (
	"errors"
	"testing"
)

func Test_stringUnpacking(t *testing.T) {
	testTable := []struct {
		innings string
		result  string
		err     error
	}{
		{
			innings: "a4bc2d5e",
			result:  "aaaabccddddde",
			err:     nil,
		},
		{
			innings: "abcd",
			result:  "abcd",
			err:     errors.New("некорректная строка"),
		},
		{
			innings: "45",
			result:  "",
			err:     errors.New("некорректная строка"),
		},
		{
			innings: "",
			result:  "",
			err:     errors.New("некорректная строка"),
		},
	}

	for _, testCase := range testTable {
		result, err := stringUnpacking(testCase.innings)
		if result != testCase.result && testCase.err != err {
			t.Errorf("Incorrect result. Expect: %s, %s, got %s, %s", testCase.result, testCase.err, result, err)
		}
	}
}
