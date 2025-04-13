package problems

// Remove Duplicates from Sorted List
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicatesFromSortedList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	c := head
	for c.Next != nil {
		v := c.Val
		if c.Next.Val == v {
			c.Next = c.Next.Next
		} else {
			c = c.Next
		}
	}
	return head
}

// Using hashmap for storing unique values
func removeDuplicatesFromUnsortedList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	seen := make(map[int]bool) // Hash map to track seen values
	seen[head.Val] = true

	current := head
	for current.Next != nil {
		if seen[current.Next.Val] {
			// Skip duplicate node
			current.Next = current.Next.Next
		} else {
			// Mark value as seen and move forward
			seen[current.Next.Val] = true
			current = current.Next
		}
	}
	return head
}
