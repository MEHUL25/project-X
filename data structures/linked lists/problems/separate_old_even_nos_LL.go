package problems

// Given a singly linked list, we need to rearrange the nodes so that:
// All even-valued nodes appear before all odd-valued nodes.
// The relative order of nodes should be preserved.

// Approach :
// Create two separate lists:
// evenHead: Stores even-valued nodes.
// oddHead: Stores odd-valued nodes.
// Traverse the list and distribute nodes into even and odd lists.
// Connect the even list to the odd list.
// Return the modified linked list.

// Definition for a ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to separate even and odd values in a linked list
func separateEvenOdd(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// Dummy nodes for even and odd lists
	evenHead, oddHead := &ListNode{}, &ListNode{}
	even, odd := evenHead, oddHead

	// Traverse the list
	curr := head
	for curr != nil {
		if curr.Val%2 == 0 { // Even node
			even.Next = curr
			even = even.Next
		} else { // Odd node
			odd.Next = curr
			odd = odd.Next
		}
		curr = curr.Next
	}

	// Connect even list to odd list and terminate odd list
	even.Next = oddHead.Next
	odd.Next = nil

	return evenHead.Next // Return the new head (even list)
}
