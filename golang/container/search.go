package container

import (
	"fmt"
	"sort"
)

func searchExample() {
	arr := []int{1, 3, 5, 7, 9, 11, 13}
	target := 7

	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	if index < len(arr) && arr[index] == target {
		fmt.Printf("Found %d at index %d\n", target, index)
	} else {
		fmt.Printf("%d not found\n", target)
	}
}

func upperLowerBoundExample() {
	arr := []int{1, 3, 3, 3, 5, 7, 9}
	target := 3

	// Lower Bound: first position where arr[i] >= target
	lower := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	// Upper Bound: first position where arr[i] > target
	upper := sort.Search(len(arr), func(i int) bool {
		return arr[i] > target
	})

	fmt.Printf("Lower bound of %d is at index %d\n", target, lower)
	fmt.Printf("Upper bound of %d is at index %d\n", target, upper)

	fmt.Printf("Count of %d = %d\n", target, upper-lower)
}
