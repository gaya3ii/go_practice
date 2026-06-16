// lc 128: Longest Consecutive Sequence
package sorting

import "sort"

// LongestConsecutive returns the longest run of consecutive integers in nums.
// e.g. [100, 200, 3, 4, 1, 2] -> [1, 2, 3, 4]
func LongestConsecutive(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })

	best := []int{sorted[0]}
	current := []int{sorted[0]}

	for i := 1; i < len(sorted); i++ {
		if sorted[i] == sorted[i-1] {
			continue // skip duplicates
		}
		if sorted[i] == sorted[i-1]+1 {
			current = append(current, sorted[i])
		} else {
			current = []int{sorted[i]}
		}
		if len(current) > len(best) {
			best = current
		}
	}

	return best
}
