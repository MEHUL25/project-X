package problems

// Use two pointers:
// Slow pointer (slow) moves one step at a time.
// Fast pointer (fast) moves two steps at a time.
// If there is a cycle, slow and fast will eventually meet.
// To find the starting node of the cycle:
// Reset one pointer (slow) to the head.
// Move both pointers one step at a time.
// The point where they meet again is the start of the cycle.

// Definition for ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectStartingPointInCycle(head *ListNode) *ListNode {
	if DetectCycle(head) {
		if head == nil || head.Next == nil {
			return nil
		}

		slow, fast := head, head

		// Step 1: Detect if a cycle exists
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next

			if slow == fast { // Cycle detected
				break
			}
		}

		// If no cycle was detected
		if fast == nil || fast.Next == nil {
			return nil
		}

		// Step 2: Find the start of the cycle
		slow = head
		for slow != fast {
			slow = slow.Next
			fast = fast.Next
		}

		return slow // The starting node of the cycle
	}
	return nil
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
