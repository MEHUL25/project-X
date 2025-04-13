package problems

// Critical points in a Linked List
// Given the head of a linked list, Find the number of critical points. (The starting and end are not considered critical points).

// Local minima or maxima are called critical points.

// A Node is called a local minima if both next and previous elements are greater than the current element.

// A Node is called a local maxima if both next and previous elements are smaller than the current element.

type ListNode struct {
	Val  int
	Next *ListNode
}

func criticalPoints(head *ListNode) int {
	prev := head
	count := 0
	if head == nil || head.Next == nil {
		return 0
	}
	head = head.Next
	for head.Next != nil {
		if prev.Val > head.Val && head.Val < head.Next.Val {
			count++
		}
		if prev.Val < head.Val && head.Val > head.Next.Val {
			count++
		}
		prev = head
		head = head.Next
	}
	return count
}
