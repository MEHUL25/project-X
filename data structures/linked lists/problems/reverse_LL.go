package problems

// Definition for a ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to reverse a linked list
func reverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next // Store next node
		curr.Next = prev  // Reverse the link
		prev = curr       // Move prev to current node
		curr = next       // Move to next node
	}
	return prev // New head of the reversed list
}
