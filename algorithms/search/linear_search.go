package search

func linearSearch(arr []int, target int) (int, bool) {
	index := -1
	found := false

	for i, _ := range arr {
		if arr[i] == target {
			index = i
			found = true
			break
		}
	}
	return index, found
}
