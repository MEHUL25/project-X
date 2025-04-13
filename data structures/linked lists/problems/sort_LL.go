package problems

// Since random access is not possible in linked lists, Merge Sort is preferred over Quick Sort.

// Find the middle of the list (using the slow & fast pointer approach).
// Recursively sort the left and right halves.
// Merge the sorted halves.

// Definition for a ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to sort a linked list using Merge Sort
func sortLinkedList(head *ListNode) *ListNode {
	// Base case: If head is nil or only one element
	if head == nil || head.Next == nil {
		return head
	}

	// Step 1: Split the list into two halves
	mid := getMiddle(head)
	nextToMid := mid.Next
	mid.Next = nil

	// Step 2: Recursively sort both halves
	left := sortLinkedList(head)
	right := sortLinkedList(nextToMid)

	// Step 3: Merge the sorted halves
	return mergeSortedLists(left, right)
}

// Function to find the middle node using slow-fast pointer method
func getMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// Function to merge two sorted linked lists
func mergeSortedLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	// Append the remaining elements
	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}

	return dummy.Next
}
