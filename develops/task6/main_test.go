package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cut(t *testing.T) {
	testTable := []struct {
		input  string
		output string
		flags  flags
	}{
		{
			input:  "Aboba\tBIBA\tBOBA\tpipipupu",
			output: "Aboba",
			flags:  flags{fields: "0", delimiter: "\t"},
		},
		{
			input:  "Aboba\tBIBA\tBOBA\tpipipupu",
			output: "BIBA",
			flags:  flags{fields: "1", delimiter: "\t"},
		},
		{
			input:  "Aboba\tBIBA\tBOBA\tpipipupu",
			output: "BIBABOBA",
			flags:  flags{fields: "1,2", delimiter: "\t"},
		},
		{
			input:  "Aboba Salfetca\tBIBA Varezhka\tBOBA Kastrulya\tpipipupu Oguretz",
			output: "AbobaSalfetca\tBIBA",
			flags:  flags{fields: "0,1", delimiter: " "},
		},
		{
			input:  "Aboba Salfetca\tBIBA Varezhka\tBOBA Kastrulya\tpipipupu Oguretz",
			output: "",
			flags:  flags{fields: "1,2", delimiter: "\n", separated: true},
		},
		{
			input:  "Aboba Salfetca\tBIBA Varezhka\tBOBA Kastrulya\tpipipupu Oguretz",
			output: "BIBA VarezhkaBOBA Kastrulya",
			flags:  flags{fields: "1,2", delimiter: "\t", separated: true},
		},
	}

	for _, testCase := range testTable {
		result := cut(testCase.input, &testCase.flags)
		assert.Equal(t, testCase.output, result)
	}
}
