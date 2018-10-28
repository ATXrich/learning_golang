package main

import "testing"

func TestValidArgs(t *testing.T) {
	argsTable := []struct {
		name   string
		input  []string
		output error
	}{
		{name: "no arguments", input: []string{"prog_name"}, output: noArgumentsError},
		{name: "valid arguments", input: []string{"foo", "bar"}, output: nil},
	}

	for _, tt := range argsTable {
		result := ValidArgs(tt.input)

		if result != tt.output {
			t.Error(
				"For", tt.name,
				"Expected", tt.output,
				"Got", result,
			)
		}
	}

}
