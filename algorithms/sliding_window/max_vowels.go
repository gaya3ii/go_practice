package sliding_window

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func MaxVowels(s string, k int) int {
	count := 0
	maxCount := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			count++
		}
	}
	maxCount = count
	for i := k; i < len(s); i++ {
		if isVowel(s[i]) {
			count++
		}

		if isVowel(s[i-k]) {
			count--
		}
		maxCount = max(maxCount, count)
	}
	return maxCount
}
