package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	seen := map[byte]int{}
	left := 0
	res := 0
	for right := 0; right < len(s); right++ {
		if lastIndex, ok := seen[s[right]]; ok {
			left = max(left, lastIndex+1)
		}
		seen[s[right]] = right
		res = max(res, right-left+1)
	}
	return res
}
func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
