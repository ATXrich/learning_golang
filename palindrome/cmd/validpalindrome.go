package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/atxrich/learning_golang/palindrome"
)

var (
	noArgumentsError = errors.New("You must provide one string argument (with no spaces)")
	//invalidArgumentsError = errors.New("Invalid argument. Please enter a string with no spaces")
)

func main() {
	err := ValidArgs(os.Args)

	if err != nil {
		fmt.Println(noArgumentsError)
	}

	if palindrome.PalindromeCheck(os.Args[1]) {
		fmt.Printf("'%s' is a palindrome!\n", os.Args[1])
	} else {
		fmt.Println("Sorry, not a palindrome!")
	}

}

func ValidArgs(args []string) error {
	if len(args) > 1 {
		return nil
	}

	return noArgumentsError
}
