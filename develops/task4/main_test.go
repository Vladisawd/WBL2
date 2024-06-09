package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sort(t *testing.T) {
	testTable := []struct {
		input  []string
		output *map[string][]string
	}{
		{
			input: []string{"Тяпка", "пятак", "лИсток", "столик", "пятка", "слиток"},
			output: &map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
		},
	}

	for _, testCase := range testTable {
		result := anagramSearch(&testCase.input)
		assert.Equal(t, testCase.output, result)
	}
}
