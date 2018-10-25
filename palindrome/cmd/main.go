package main

import (
	"errors"
)

var (
	noArgumentsError      = errors.New("You must provide one string argument (with no spaces)")
	invalidArgumentsError = errors.New("Too many arguments. Please enter a string with no spaces")
)

func main() {

	// palindrome.PalindromeCheck(os.Args)

}

func ArgsCheck(input []string) (string, error) {
	switch {
	case len(input) == 0:
		return "", noArgumentsError
		//os.Exit(1)
	case len(input) > 1:
		return "", invalidArgumentsError
		//os.Exit(1)
	}
	return input[0], nil
}
