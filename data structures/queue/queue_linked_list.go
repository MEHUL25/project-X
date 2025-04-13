package queue

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type Queue struct {
	front *Node
	rear  *Node
	size  int
}

// Enqueue operation to add an element to the queue
func (q *Queue) Enqueue(value int) {
	newNode := &Node{data: value}
	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
	q.size++
}

// Dequeue operation to remove and return the front element of the queue
func (q *Queue) Dequeue() (int, error) {
	if q.front == nil {
		return 0, fmt.Errorf("queue is empty")
	}
	frontValue := q.front.data
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	q.size--
	return frontValue, nil
}

// Peek operation to return the front element without removing it
func (q *Queue) Peek() (int, error) {
	if q.front == nil {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.front.data, nil
}

// IsEmpty operation to check if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.front == nil
}

// Size operation to return the number of elements in the queue
func (q *Queue) Size() int {
	return q.size
}

func queueLinkedListVersion() {
	queue := &Queue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	fmt.Println("Queue after enqueuing elements:")
	temp := queue.front
	for temp != nil {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.next
	}
	fmt.Println("NULL")

	front, _ := queue.Dequeue()
	fmt.Println("Dequeued element:", front)

	peek, _ := queue.Peek()
	fmt.Println("Front element (peek):", peek)
	fmt.Println("Is queue empty?", queue.IsEmpty())
	fmt.Println("Queue size:", queue.Size())
}
