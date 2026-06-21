package prefix_sum

func NumberOfSubarrays(nums []int, k int) int {
	oddCount := 0
	freqCount := map[int]int{0: 1}
	count := 0
	for _, num := range nums {
		oddCount += num % 2
		count += freqCount[oddCount-k]
		freqCount[oddCount]++
	}
	return count
}
