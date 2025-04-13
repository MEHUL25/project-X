package list

import "fmt"

type Node struct {
	data int
	next *Node
}

type CircularLinkedList struct {
	head *Node
}

// Delete from a specific position in the circular linked list
func (cll *CircularLinkedList) DeleteFromK(k int) {
	if k <= 0 {
		fmt.Println("Invalid position")
		return
	}

	if cll.head == nil {
		fmt.Println("List is empty")
		return
	}

	if k == 1 {
		temp := cll.head
		if temp.next == cll.head {
			cll.head = nil
			return
		}
		for temp.next != cll.head {
			temp = temp.next
		}
		cll.head = cll.head.next
		temp.next = cll.head
		return
	}

	temp := cll.head
	for i := 1; i < k-1 && temp.next != cll.head; i++ {
		temp = temp.next
	}

	if temp.next == cll.head {
		fmt.Println("Position out of bounds")
		return
	}

	temp.next = temp.next.next
}
