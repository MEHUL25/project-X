package search

func binarySearchIterative(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	mid := 0

	for low >= 0 && high <= len(arr)-1 && low <= high {
		mid = (low + high) / 2
		if target == arr[mid] {
			return mid
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
