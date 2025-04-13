package problems

// The intersection of two linked lists occurs when they share a common node (by reference, not by value). Our goal is to find the node where they intersect.

// Traverse both lists to find their lengths.
// Align the heads of both lists by skipping extra nodes in the longer list.
// Move both pointers simultaneously until they meet at the intersection point (or reach nil if no intersection exists).

// Definition for a ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to find the intersection node of two linked lists
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := getLength(headA), getLength(headB)

	// Align the start position of both lists
	for lenA > lenB {
		headA = headA.Next
		lenA--
	}
	for lenB > lenA {
		headB = headB.Next
		lenB--
	}

	// Move both pointers together until they meet
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA // Returns the intersection node or nil if no intersection
}

// Function to calculate the length of a linked list
func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}
