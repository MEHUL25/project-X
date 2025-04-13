package problems

// Definition for ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to find the middle of the linked list
func findMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	// Move fast two steps and slow one step at a time
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Slow pointer is now at the middle of the list
	return slow
}
