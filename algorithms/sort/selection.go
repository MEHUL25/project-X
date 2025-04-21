package sort

// Divides the input array into two parts: a sorted part and an unsorted part.
// Repeatedly finds the minimum element in the unsorted part and swaps it with the leftmost element of the unsorted part.
// Expands the sorted part by one element in each iteration.

func selectionSort(arr []int) []int {
	pos := 0
	for i := 0; i < len(arr)-1; i++ {
		pos = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[pos] {
				pos = j
			}
		}
		if pos != i {
			arr[i], arr[pos] = arr[pos], arr[i]
		}
	}
	return arr
}
