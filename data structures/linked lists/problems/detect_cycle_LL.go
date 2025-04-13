package problems

// Floyd’s Cycle Detection Algorithm (also known as the Tortoise and Hare algorithm)

// Definition for ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to detect the start of cycle in a linked list
func DetectCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head

	// Step 1: Detect if a cycle exists
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast { // Cycle detected
			return true
		}
	}

	// If no cycle was detected
	if fast == nil || fast.Next == nil {
		return false
	}
	return false
}
