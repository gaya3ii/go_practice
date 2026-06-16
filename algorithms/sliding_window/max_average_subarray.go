// lc 643: Maximum Average Subarray I
package sliding_window

func FindMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := range k {
		sum += nums[i]
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i]
		sum -= nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return float64(maxSum) / float64(k)
}
