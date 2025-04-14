package heap

// Comparison Logic: Change the Less method to prioritize larger values (i.e., we want to swap elements when one has a higher priority than the other).

// Item defines the structure of the elements in the priority queue
type Item struct {
	value    string // The value of the item (can be anything)
	priority int    // The priority (lower = higher priority)
	index    int    // The index of the item in the heap
}

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

// Less returns true if the item at index i has a higher priority than the item at index j (max-heap).
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority // Max-heap (reverse the comparison)
}
