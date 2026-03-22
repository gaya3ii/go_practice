package main

import (
	"fmt"
	"unicode"
)

func isAlnum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isAlnum(rune(s[left])) {
			left++
		}

		for left < right && !isAlnum(rune(s[right])) {
			right--
		}

		l := unicode.ToLower(rune(s[left]))
		r := unicode.ToLower(rune(s[right]))

		if l != r {
			return false
		}

		left++
		right--
	}

	return true
}

func main() {
	s := "racecar"
	fmt.Println(isPalindrome(s))
}
