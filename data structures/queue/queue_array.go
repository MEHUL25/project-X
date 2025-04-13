package queue

import (
	"fmt"
)

type Queue struct {
	elements []int
}

// Enqueue operation to add an element to the queue
func (q *Queue) Enqueue(value int) {
	q.elements = append(q.elements, value)
}

// Dequeue operation to remove and return the front element of the queue
func (q *Queue) Dequeue() (int, error) {
	if len(q.elements) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	front := q.elements[0]
	q.elements = q.elements[1:]
	return front, nil
}

// Peek operation to return the front element without removing it
func (q *Queue) Peek() (int, error) {
	if len(q.elements) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.elements[0], nil
}

// IsEmpty operation to check if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

// Size operation to return the number of elements in the queue
func (q *Queue) Size() int {
	return len(q.elements)
}

func queueArrayVersion() {
	queue := &Queue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	fmt.Println("Queue after enqueuing elements:", queue.elements)

	front, _ := queue.Dequeue()
	fmt.Println("Dequeued element:", front)
	fmt.Println("Queue after dequeuing:", queue.elements)

	peek, _ := queue.Peek()
	fmt.Println("Front element (peek):", peek)
	fmt.Println("Is queue empty?", queue.IsEmpty())
	fmt.Println("Queue size:", queue.Size())
}
