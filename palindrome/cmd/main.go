package main

import (
	"errors"
	"os"

	"github.com/atxrich/learning_golang/palindrome"
)

var (
	noArgumentsError      = errors.New("You must provide one string argument (with no spaces)")
	invalidArgumentsError = errors.New("Too many arguments. Please enter a string with no spaces")
)

func main() {

	word, err := ArgsCheck(os.Args)

	if err != nil {
		os.Exit(1)
	}

	palindrome.PalindromeCheck(word)

}

func ArgsCheck(input []string) (string, error) {
	if len(input) == 0 {
		return "", noArgumentsError
	} else if len(input) > 1 {
		return "", invalidArgumentsError
	}

	return input[0], nil
}
