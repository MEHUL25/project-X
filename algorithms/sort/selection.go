package sort

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
