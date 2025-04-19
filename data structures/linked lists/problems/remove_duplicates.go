package problems

// Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

// Traverse the list

// Compare each node with the next

// Skip the next if it’s a duplicate

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			// Skip the duplicate
			current.Next = current.Next.Next
		} else {
			// Move to the next distinct element
			current = current.Next
		}
	}

	return head
}
