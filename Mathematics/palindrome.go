package Mathematics

func Palindrome(s string) bool {
	// checks for palindrome by reversing the sting
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return s == string(runes)
}
