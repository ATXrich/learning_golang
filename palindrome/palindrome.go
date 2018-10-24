package palindrome

func PalindromeCheck(word string) bool {
	chars := []rune(word)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return word == string(chars)

}
