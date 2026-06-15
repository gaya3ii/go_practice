package main

func longestSubarray(nums []int) int {
	zero := 0
	left := 0
	maxlength := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zero += 1
		}
		for zero > 1 {
			if nums[left] == 0 {
				zero--
			}
			left += 1
		}
		maxlength = max(maxlength, right-left+1)
	}
	return maxlength
}

func main() {
	nums := []int{1, 1, 0, 1}
	println(longestSubarray(nums))
}
