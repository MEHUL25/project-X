package search

func binarySearchRecursive(arr []int, target int, low, high int) int {

	if low <= high {

		mid := (low + high) / 2
		if target == arr[mid] {
			return mid
		} else if target < arr[mid] {
			return binarySearchRecursive(arr, target, low, mid-1)
		} else {
			return binarySearchRecursive(arr, target, mid+1, high)
		}
	}
	return -1
}
