package container

import (
	"container/heap"
	"fmt"
	"sort"
)

// Heap struct defined in library
type Interface interface {
	sort.Interface      // Len(), Less(i, j), Swap(i, j)
	Push(x interface{}) // Push element onto heap
	Pop() interface{}   // Remove and return min/max
}

// Function	Description
// heap.Init(h)	Initializes a heap from an array
// heap.Push(h, x)	Adds element to the heap
// heap.Pop(h)	Removes and returns the top (min/max)
// heap.Fix(h, i)	Re-establishes heap order after change
// heap.Remove(h, i)	Removes the element at index i

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] } // Min-heap
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func heapExample() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Println(heap.Pop(h)) // 1 (min)
}
