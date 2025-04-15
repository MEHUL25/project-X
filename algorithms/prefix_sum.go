package algorithms

// Given an array: nums = [1, 2, 3, 4]
// We create a prefix sum array: prefix = [0, 1, 3, 6, 10]
// Then:
// Sum of nums[1:3] → prefix[3] - prefix[1] = 6 - 1 = 5
// prefix[0] = 0 for easy calculation (common convention).
// Efficient for multiple range sum queries (O(1) after O(n) preprocessing).

import (
	"fmt"
)

// BuildPrefixSum constructs the prefix sum array
func BuildPrefixSum(nums []int) []int {
	n := len(nums)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	return prefix
}

// RangeSum returns the sum of nums[left:right+1] using the prefix sum array
func RangeSum(prefix []int, left, right int) int {
	return prefix[right+1] - prefix[left]
}

func example1() {
	nums := []int{1, 2, 3, 4, 5}
	prefix := BuildPrefixSum(nums)

	fmt.Println("Prefix Sum:", prefix)
	fmt.Println("Sum from index 1 to 3:", RangeSum(prefix, 1, 3)) // 2+3+4 = 9
	fmt.Println("Sum from index 0 to 4:", RangeSum(prefix, 0, 4)) // 1+2+3+4+5 = 15
}
