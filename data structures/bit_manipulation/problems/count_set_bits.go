package problems

// Count the number of 1s in binary representation of a number.
func countSetBits(n int) int {
	count := 0
	for n > 0 {
		n = n & (n - 1)
		count++
	}
	return count
}
