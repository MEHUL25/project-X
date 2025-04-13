package problems

// Rotating a singly linked list by k positions means shifting the nodes so that the last k nodes move to the front while maintaining their relative order.

// Find the length (N) of the list.
// Normalize k (if k >= N, then k = k % N).
// Find the new tail (N-k-1 th node) and break the list.
// Connect the old tail to the head and update the new head.

// Definition for ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to rotate the linked list by k places
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// Step 1: Find the length of the list
	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	// Step 2: Normalize k
	k = k % length
	if k == 0 {
		return head // No rotation needed
	}

	// Step 3: Find the new tail (N-k-1) and new head (N-k)
	newTail := head
	for i := 0; i < length-k-1; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next

	// Step 4: Break the list and connect the tail to the original head
	newTail.Next = nil
	tail.Next = head

	return newHead
}
