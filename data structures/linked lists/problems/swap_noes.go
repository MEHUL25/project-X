package problems

// Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

// We'll use a dummy node to simplify edge cases (like swapping the first pair).

// Traverse the list two nodes at a time, swap them, and reconnect the links properly.
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	// Dummy node to handle edge cases easily
	dummy := &ListNode{Next: head}
	current := dummy

	for current.Next != nil && current.Next.Next != nil {
		// Nodes to be swapped
		first := current.Next
		second := current.Next.Next

		// Swapping
		first.Next = second.Next
		second.Next = first
		current.Next = second

		// Move to the next pair
		current = first
	}

	return dummy.Next
}
