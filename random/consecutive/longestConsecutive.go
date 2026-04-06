package main

import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	return nums
}
func main() {
	k := longestConsecutive([]int{100, 200, 3, 4, 1, 2})
	fmt.Println(k)
}
