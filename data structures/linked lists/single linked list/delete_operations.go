package list

import "fmt"

type LinkedList struct {
	head *Node
}

type Node struct {
	data int
	next *Node
}

// Delete from the front of the linked list
func (ll *LinkedList) DeleteFromFront() {
	if ll.head == nil {
		fmt.Println("List is empty")
		return
	}
	ll.head = ll.head.next
}

// Delete from a specific position in the linked list
func (ll *LinkedList) DeleteFromPosition(k int) {
	if k <= 0 {
		fmt.Println("Invalid position")
		return
	}

	if ll.head == nil {
		fmt.Println("List is empty")
		return
	}

	if k == 1 {
		ll.head = ll.head.next
		return
	}

	temp := ll.head
	for i := 1; i < k-1 && temp.next != nil; i++ {
		temp = temp.next
	}

	if temp.next == nil {
		fmt.Println("Position out of bounds")
		return
	}

	temp.next = temp.next.next
}
