// lc 3: Longest Substring Without Repeating Characters
package sliding_window

func LengthOfLongestSubstring(s string) int {
	seen := map[byte]int{}
	left := 0
	answer := 0

	for right := 0; right < len(s); right++ {
		// Step 1: read OLD position (before updating)
		if lastIdx, ok := seen[s[right]]; ok && lastIdx >= left {
			left = lastIdx + 1
		}

		// Step 2: update to NEW position
		seen[s[right]] = right

		// Step 3: update answer
		answer = max(answer, right-left+1)
	}

	return answer
}
