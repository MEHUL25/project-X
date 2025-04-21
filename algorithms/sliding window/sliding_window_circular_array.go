package algorithms

import "fmt"

// Idea: Treat the array as doubled (2n) by using i % n

func circularSlidingWindow(nums []int, windowSize int) {
	n := len(nums)
	left := 0
	currentSum := 0

	for right := 0; right < 2*n; right++ {
		idx := right % n
		currentSum += nums[idx]

		// If window exceeds desired size, shrink from left
		if right-left+1 > windowSize {
			currentSum -= nums[left%n]
			left++
		}

		// If window hits desired size, do something
		if right-left+1 == windowSize {
			fmt.Println("Window sum:", currentSum)
		}
	}
}
