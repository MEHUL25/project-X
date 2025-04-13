package sort

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
