package prefix_sum

func NumberOfSubarrays(nums []int, k int) int {
	OddCount := 0
	freqCount := map[int]int{0: 1}
	count := 0
	for _, num := range nums {
		OddCount += num % 2
		count += freqCount[OddCount-k]
		freqCount[OddCount]++
	}
	return count
}
