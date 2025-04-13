package problems

// In an array where every number appears twice except one, find that one.

// Trick: XOR cancels duplicates.

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}
