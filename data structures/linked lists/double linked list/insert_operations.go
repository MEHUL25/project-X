package list

import (
	"fmt"
)

type Node struct {
	data int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
}

// Insert at the k-th position in the doubly linked list (1-based index)
func (dll *DoublyLinkedList) InsertAtK(data, k int) {
	newNode := &Node{data: data}
	if k <= 0 {
		fmt.Println("Invalid position")
		return
	}

	if dll.head == nil {
		if k == 1 {
			dll.head = newNode
		} else {
			fmt.Println("Position out of bounds")
		}
		return
	}

	if k == 1 {
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
		return
	}

	temp := dll.head
	for i := 1; i < k-1 && temp.next != nil; i++ {
		temp = temp.next
	}

	if temp.next == nil && k > 1 {
		temp.next = newNode
		newNode.prev = temp
	} else {
		newNode.next = temp.next
		newNode.prev = temp
		if temp.next != nil {
			temp.next.prev = newNode
		}
		temp.next = newNode
	}
}

// Print the doubly linked list
func (dll *DoublyLinkedList) PrintList() {
	temp := dll.head
	for temp != nil {
		fmt.Printf("%d <-> ", temp.data)
		temp = temp.next
	}
	fmt.Println("NULL")
}

func doublyLinkedListExample() {
	dll := &DoublyLinkedList{}

	dll.InsertAtK(10, 1)
	dll.InsertAtK(20, 2)
	dll.InsertAtK(30, 3)
	fmt.Print("After inserting at positions: ")
	dll.PrintList()

	// dll.DeleteFromK(2)
	// fmt.Print("After deleting from 2nd position: ")
	// dll.PrintList()
}
