package problems

// You are given the head of a non-empty linked list representing a non-negative integer without leading zeroes.

// Return the head of the linked list after doubling it.

// Key steps:
// Use a recursive function to reach the last node.

// On returning back up:

// Double the current node value.

// Add any carry from deeper calls.

// Update current node value.

// Return carry up to previous node.

// If there's a final carry after the full recursion, insert a new node at the head.

type ListNode struct {
	Val  int
	Next *ListNode
}

func doubleIt(head *ListNode) *ListNode {
	// Helper recursive function
	var dfs func(node *ListNode) int
	dfs = func(node *ListNode) int {
		if node == nil {
			return 0
		}
		carry := dfs(node.Next)
		sum := node.Val*2 + carry
		node.Val = sum % 10
		return sum / 10
	}

	carry := dfs(head)
	if carry > 0 {
		newHead := &ListNode{Val: carry, Next: head}
		return newHead
	}
	return head
}
