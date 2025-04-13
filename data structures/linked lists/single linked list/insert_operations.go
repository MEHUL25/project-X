package list

import (
	"fmt"
)

// Insert At Front
// Insert At end
// Insert At K th position

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// Insert at the end of the linked list
func (ll *LinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	temp := ll.head

	// Find the last node, So temp doesn't become nil, it stops before hand only
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

// Insert at the front of the linked list
func (ll *LinkedList) InsertAtFront(data int) {
	newNode := &Node{data: data, next: ll.head}
	ll.head = newNode
}

// Insert at the k-th position in the linked list (1-based index)
func (ll *LinkedList) InsertAtK(data, k int) {
	if k <= 0 {
		fmt.Println("Invalid position")
		return
	}

	newNode := &Node{data: data}
	if k == 1 {
		newNode.next = ll.head
		ll.head = newNode
		return
	}

	temp := ll.head
	for i := 1; i < k-1 && temp != nil; i++ {
		temp = temp.next
	}

	if temp == nil {
		fmt.Println("Position out of bounds")
		return
	}

	newNode.next = temp.next
	temp.next = newNode
}

// Print the linked list
func (ll *LinkedList) PrintList() {
	temp := ll.head
	for temp != nil {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.next
	}
	fmt.Println("NULL")
}

func main() {
	ll := &LinkedList{}

	ll.InsertAtEnd(10)
	ll.InsertAtEnd(20)
	ll.InsertAtEnd(30)
	fmt.Print("After inserting at end: ")
	ll.PrintList()

	ll.InsertAtFront(5)
	fmt.Print("After inserting at front: ")
	ll.PrintList()

	ll.InsertAtK(15, 3)
	fmt.Print("After inserting at 3rd position: ")
	ll.PrintList()
}
