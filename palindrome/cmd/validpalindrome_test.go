package main

import "testing"

func TestArgsCheck(t *testing.T) {
	argsTable := []struct {
		input  []string
		output string
		err    error
	}{
		{input: []string{"foo"}, output: "foo", err: nil},
		{input: []string{}, output: "", err: noArgumentsError},
		{input: []string{"foo", "bar"}, output: "", err: invalidArgumentsError},
	}

	for _, tt := range argsTable {
		result, err := ArgsCheck(tt.input)

		if result != tt.output || err != tt.err {
			t.Error(
				"Want:", tt.output, tt.err,
				"Got:", result, err,
			)
		}
	}

}
