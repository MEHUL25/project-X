package algorithms

func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	maxLen := 0

	for right < len(nums) {
		// If it's a zero, we consume a flip
		if nums[right] == 0 {
			k--
		}

		// If flips used exceed k, move left pointer forward
		for k < 0 {
			if nums[left] == 0 {
				k++
			}
			left++
		}

		// Update max length
		if (right - left + 1) > maxLen {
			maxLen = right - left + 1
		}
		right++
	}

	return maxLen
}
