package problems

// Definition for ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to count loop length in a linked list
func countLoopLength(head *ListNode) int {
	slow, fast := head, head

	// Step 1: Detect the cycle using Floyd's Algorithm
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// If slow and fast meet, a loop is detected
		if slow == fast {
			// Step 2: Count loop length
			count := 1
			current := slow.Next
			for current != slow {
				count++
				current = current.Next
			}
			return count
		}
	}

	// No loop found
	return 0
}
