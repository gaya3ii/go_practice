// lc 125: Valid Palindrome
package two_pointers

import "unicode"

func isAlnum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func IsPalindrome(s string) bool {
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
