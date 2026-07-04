package sliding_window

func numOfSubarrays(arr []int, k int, threshold int) int {

	sum := 0

	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	count := 0
	if sum >= threshold*k {
		count++
	}

	for i := k; i < len(arr); i++ {
		sum += arr[i]
		sum -= arr[i-k]
		if sum >= threshold*k {
			count++
		}
	}
	return count

}
