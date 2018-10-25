package main

import "testing"

func TestArgsCheck(t *testing.T) {
	argsTable := []struct {
		input  []string
		output string
		//err    error
	}{
		{input: []string{"foo"}, output: "foo"},
		{input: []string{}, output: ""},
		{input: []string{"foo", "bar"}, output: ""},
	}

	for _, tt := range argsTable {
		result := ArgsCheck(tt.input)

		if result != tt.output {
			t.Error(
				"Want:", tt.output,
				"Got:", result,
			)
		}
	}

}
