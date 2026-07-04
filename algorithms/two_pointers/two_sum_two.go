package two_pointers

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		current := numbers[left] + numbers[right]

		if current == target {
			return []int{left + 1, right + 1}
		} else if current < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
