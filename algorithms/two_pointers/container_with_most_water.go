package two_pointers

func container_with_most_water(arr []int) int {
	left, right := 0, len(arr)-1
	best := 0
	for left < right {
		current := compute(arr[left], arr[right], right-left)
		best = max(best, current)
		if arr[left] < arr[right] {
			left++
		} else {
			right--
		}
	}
	return best
}

func compute(h1, h2, width int) int {
	return min(h1, h2) * width
}
