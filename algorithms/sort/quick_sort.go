package sort

// QuickSort sorts the slice in-place using the QuickSort algorithm
func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// Partitioning
	pivotIndex := partition(arr)

	// Recursively sort left and right partitions
	QuickSort(arr[:pivotIndex])
	QuickSort(arr[pivotIndex+1:])
}

// Partition function: rearranges the slice and returns the pivot index
func partition(arr []int) int {
	pivot := arr[len(arr)-1]
	i := 0

	for j := 0; j < len(arr)-1; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// Swap pivot to the correct position
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	return i
}
