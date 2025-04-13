package list

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
}

// Delete from a specific position in the doubly linked list
func (dll *DoublyLinkedList) DeleteFromK(k int) {
	if k <= 0 {
		fmt.Println("Invalid position")
		return
	}

	if dll.head == nil {
		fmt.Println("List is empty")
		return
	}

	if k == 1 {
		dll.head = dll.head.next
		if dll.head != nil {
			dll.head.prev = nil
		}
		return
	}

	temp := dll.head
	for i := 1; i < k && temp != nil; i++ {
		temp = temp.next
	}

	if temp == nil {
		fmt.Println("Position out of bounds")
		return
	}

	if temp.next != nil {
		temp.next.prev = temp.prev
	}
	if temp.prev != nil {
		temp.prev.next = temp.next
	}
}
