package queue

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

type Deque struct {
	front *Node
	rear  *Node
	size  int
}

// EnqueueFront operation to add an element to the front of the deque
func (d *Deque) EnqueueFront(value int) {
	newNode := &Node{data: value}
	if d.front == nil {
		d.front = newNode
		d.rear = newNode
	} else {
		newNode.next = d.front
		d.front.prev = newNode
		d.front = newNode
	}
	d.size++
}

// EnqueueRear operation to add an element to the rear of the deque
func (d *Deque) EnqueueRear(value int) {
	newNode := &Node{data: value}
	if d.rear == nil {
		d.front = newNode
		d.rear = newNode
	} else {
		d.rear.next = newNode
		newNode.prev = d.rear
		d.rear = newNode
	}
	d.size++
}

// DequeueFront operation to remove and return the front element of the deque
func (d *Deque) DequeueFront() (int, error) {
	if d.front == nil {
		return 0, fmt.Errorf("deque is empty")
	}
	value := d.front.data
	d.front = d.front.next
	if d.front == nil {
		d.rear = nil
	} else {
		d.front.prev = nil
	}
	d.size--
	return value, nil
}

// DequeueRear operation to remove and return the rear element of the deque
func (d *Deque) DequeueRear() (int, error) {
	if d.rear == nil {
		return 0, fmt.Errorf("deque is empty")
	}
	value := d.rear.data
	d.rear = d.rear.prev
	if d.rear == nil {
		d.front = nil
	} else {
		d.rear.next = nil
	}
	d.size--
	return value, nil
}

// PeekFront operation to return the front element without removing it
func (d *Deque) PeekFront() (int, error) {
	if d.front == nil {
		return 0, fmt.Errorf("deque is empty")
	}
	return d.front.data, nil
}

// PeekRear operation to return the rear element without removing it
func (d *Deque) PeekRear() (int, error) {
	if d.rear == nil {
		return 0, fmt.Errorf("deque is empty")
	}
	return d.rear.data, nil
}

// IsEmpty operation to check if the deque is empty
func (d *Deque) IsEmpty() bool {
	return d.front == nil
}

// Size operation to return the number of elements in the deque
func (d *Deque) Size() int {
	return d.size
}

func dequeueLinkedListExample() {
	deque := &Deque{}

	deque.EnqueueRear(10)
	deque.EnqueueRear(20)
	deque.EnqueueFront(5)
	fmt.Println("Deque after enqueuing elements:")
	temp := deque.front
	for temp != nil {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.next
	}
	fmt.Println("NULL")

	front, _ := deque.DequeueFront()
	fmt.Println("Dequeued from front:", front)

	rear, _ := deque.DequeueRear()
	fmt.Println("Dequeued from rear:", rear)

	peekFront, _ := deque.PeekFront()
	fmt.Println("Front element (peek):", peekFront)
	peekRear, _ := deque.PeekRear()
	fmt.Println("Rear element (peek):", peekRear)
	fmt.Println("Is deque empty?", deque.IsEmpty())
	fmt.Println("Deque size:", deque.Size())
}
