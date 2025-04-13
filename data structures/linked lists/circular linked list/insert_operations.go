package list

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type CircularLinkedList struct {
	head *Node
}

// Insert at the k-th position in the circular linked list (1-based index)
func (cll *CircularLinkedList) InsertAtK(data, k int) {
	newNode := &Node{data: data}
	if k <= 0 {
		fmt.Println("Invalid position")
		return
	}

	if cll.head == nil {
		if k == 1 {
			newNode.next = newNode
			cll.head = newNode
		} else {
			fmt.Println("Position out of bounds")
		}
		return
	}

	if k == 1 {
		temp := cll.head
		for temp.next != cll.head {
			temp = temp.next
		}
		newNode.next = cll.head
		temp.next = newNode
		cll.head = newNode
		return
	}

	temp := cll.head
	for i := 1; i < k-1 && temp.next != cll.head; i++ {
		temp = temp.next
	}

	if temp.next == cll.head && k > 1 {
		fmt.Println("Position out of bounds")
		return
	}

	newNode.next = temp.next
	temp.next = newNode
}

// Print the circular linked list
func (cll *CircularLinkedList) PrintList() {
	if cll.head == nil {
		fmt.Println("List is empty")
		return
	}
	temp := cll.head
	for {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.next
		if temp == cll.head {
			break
		}
	}
	fmt.Println("(head)")
}

func circularLinkedListExample() {
	cll := &CircularLinkedList{}

	cll.InsertAtK(10, 1)
	cll.InsertAtK(20, 2)
	cll.InsertAtK(30, 3)
	fmt.Print("After inserting at positions: ")
	cll.PrintList()

	// cll.DeleteFromK(2)
	// fmt.Print("After deleting from 2nd position: ")
	// cll.PrintList()
}
