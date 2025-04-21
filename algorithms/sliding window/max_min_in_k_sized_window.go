package algorithms

import "container/list"

// Idea: Keep a deque with indices of useful elements in decreasing (or increasing) order
// Example: maximum in each k-sized window

func maxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	deque := list.New()

	for i := 0; i < len(nums); i++ {
		// Remove indices out of the current window
		if deque.Len() > 0 && deque.Front().Value.(int) <= i-k {
			deque.Remove(deque.Front())
		}

		// Remove elements smaller than current num from the back
		for deque.Len() > 0 && nums[deque.Back().Value.(int)] < nums[i] {
			deque.Remove(deque.Back())
		}

		// Add current index
		deque.PushBack(i)

		// Add max in current window to result
		if i >= k-1 {
			result = append(result, nums[deque.Front().Value.(int)])
		}
	}
	return result
}
