package sort

// MergeSort sorts the slice using the Merge Sort algorithm
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Find the middle point and divide the array
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// Merge the two sorted halves
	return merge(left, right)
}

// Merge function to combine two sorted slices
func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	// Merge the two sorted arrays
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
