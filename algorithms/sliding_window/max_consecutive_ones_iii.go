// lc 1004: Max Consecutive Ones III
package sliding_window

func LongestOnes(nums []int, k int) int {
	left := 0
	zeros := 0
	maxLen := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros++
		}
		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}
