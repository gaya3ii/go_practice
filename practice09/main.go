package practice09

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
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

func main() {

	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	result := findMaxAverage(nums, k)
	fmt.Printf("The maximum average of a subarray of length %d is: %.5f\n", k, result)

}
