package algorithms

// Idea: Use a map to store prefix sums and their counts to efficiently track subarrays with desired sum property.
// Example: number of subarrays with sum == target

func subarraySum(nums []int, target int) int {
	count := 0
	prefixSum := 0
	sumMap := map[int]int{0: 1} // prefix sum 0 occurs once initially

	for _, num := range nums {
		prefixSum += num

		if sumMap[prefixSum-target] > 0 {
			count += sumMap[prefixSum-target]
		}

		sumMap[prefixSum]++
	}
	return count
}
