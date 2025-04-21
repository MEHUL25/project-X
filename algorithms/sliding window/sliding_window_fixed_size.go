package algorithms

import (
	"fmt"
)

// Let’s say you want the maximum sum of any subarray of size k.

func maxSumSlidingWindow(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}

	// Calculate initial window sum
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}

	maxSum := windowSum

	// Slide the window
	for i := k; i < len(nums); i++ {
		windowSum += nums[i] - nums[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 1}
	k := 3
	fmt.Println("Max sum of any subarray of size", k, "is:", maxSumSlidingWindow(nums, k))
}
