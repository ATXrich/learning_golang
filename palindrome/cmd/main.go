package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	noArgumentsError      = errors.New("You must provide one string argument (with no spaces)")
	invalidArgumentsError = errors.New("Too many arguments. Please enter a string with no spaces")
)

func main() {

	// palindrome.PalindromeCheck(os.Args)

}

func ArgsCheck(input []string) string {
	switch {
	case len(input) == 0:
		fmt.Println("no args")
		os.Exit(1)
	case len(input) > 1:
		fmt.Println("too many args")
		os.Exit(1)
		// default:
		// 	return input[0]
	}
	return input[0]
}
