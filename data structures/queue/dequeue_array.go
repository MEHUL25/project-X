package queue

import (
	"fmt"
)

type Deque struct {
	elements []int
}

// EnqueueFront operation to add an element to the front of the deque
func (d *Deque) EnqueueFront(value int) {
	d.elements = append([]int{value}, d.elements...)
}

// EnqueueRear operation to add an element to the rear of the deque
func (d *Deque) EnqueueRear(value int) {
	d.elements = append(d.elements, value)
}

// DequeueFront operation to remove and return the front element of the deque
func (d *Deque) DequeueFront() (int, error) {
	if len(d.elements) == 0 {
		return 0, fmt.Errorf("deque is empty")
	}
	front := d.elements[0]
	d.elements = d.elements[1:]
	return front, nil
}

// DequeueRear operation to remove and return the rear element of the deque
func (d *Deque) DequeueRear() (int, error) {
	if len(d.elements) == 0 {
		return 0, fmt.Errorf("deque is empty")
	}
	rear := d.elements[len(d.elements)-1]
	d.elements = d.elements[:len(d.elements)-1]
	return rear, nil
}

// PeekFront operation to return the front element without removing it
func (d *Deque) PeekFront() (int, error) {
	if len(d.elements) == 0 {
		return 0, fmt.Errorf("deque is empty")
	}
	return d.elements[0], nil
}

// PeekRear operation to return the rear element without removing it
func (d *Deque) PeekRear() (int, error) {
	if len(d.elements) == 0 {
		return 0, fmt.Errorf("deque is empty")
	}
	return d.elements[len(d.elements)-1], nil
}

// IsEmpty operation to check if the deque is empty
func (d *Deque) IsEmpty() bool {
	return len(d.elements) == 0
}

// Size operation to return the number of elements in the deque
func (d *Deque) Size() int {
	return len(d.elements)
}

func dequeueArrayExample() {
	deque := &Deque{}

	deque.EnqueueRear(10)
	deque.EnqueueRear(20)
	deque.EnqueueFront(5)
	fmt.Println("Deque after enqueuing elements:", deque.elements)

	front, _ := deque.DequeueFront()
	fmt.Println("Dequeued from front:", front)
	fmt.Println("Deque after dequeuing front:", deque.elements)

	rear, _ := deque.DequeueRear()
	fmt.Println("Dequeued from rear:", rear)
	fmt.Println("Deque after dequeuing rear:", deque.elements)

	peekFront, _ := deque.PeekFront()
	fmt.Println("Front element (peek):", peekFront)
	peekRear, _ := deque.PeekRear()
	fmt.Println("Rear element (peek):", peekRear)
	fmt.Println("Is deque empty?", deque.IsEmpty())
	fmt.Println("Deque size:", deque.Size())
}
