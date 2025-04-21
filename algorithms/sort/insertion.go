package sort

// We start with the second element of the array as the first element is assumed to be sorted.
// Compare the second element with the first element if the second element is smaller then swap them.
// Move to the third element, compare it with the first two elements, and put it in its correct position
// Repeat until the entire array is sorted.

func insertionSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		j := i - 1
		key := arr[i]
		for ; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
	return arr
}
