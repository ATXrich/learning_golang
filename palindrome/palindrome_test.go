package palindrome

import "testing"

func TestPalindromeCheck(t *testing.T) {
	inputTable := []struct {
		input  string
		output bool
	}{
		{"kayak", true},
		{"dog", false},
		{"madam", true},
		{"racecar", true},
		{"frog", false},
	}

	for _, tt := range inputTable {
		result := PalindromeCheck(tt.input)
		if result != tt.output {
			t.Error(
				"For", tt.input,
				"expected", tt.output,
				"got", result,
			)
		}
	}

}
