package queue

import (
	"container/heap"
	"fmt"
)

// Item defines the structure of the elements in the priority queue
type Item struct {
	value    string // The value of the item (can be anything)
	priority int    // The priority (lower = higher priority)
	index    int    // The index of the item in the heap
}

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority // Min-heap
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func priorityQueueExample() {
	items := map[string]int{
		"banana": 3,
		"apple":  2,
		"pear":   4,
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Push items
	for value, priority := range items {
		heap.Push(&pq, &Item{
			value:    value,
			priority: priority,
		})
	}

	// Add another item
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)

	// Update an item (optional)
	pq.update(item, item.value, 5)

	// Pop items by priority
	fmt.Println("Items in priority order:")
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%s (priority: %d)\n", item.value, item.priority)
	}
}

// Output :

// Items in priority order:
// apple (priority: 2)
// banana (priority: 3)
// pear (priority: 4)
// orange (priority: 5)
