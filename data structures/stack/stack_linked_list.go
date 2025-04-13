package stack

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top  *Node
	size int
}

// Push operation to add an element to the stack
func (s *Stack) Push(value int) {
	newNode := &Node{data: value, next: s.top}
	s.top = newNode
	s.size++
}

// Pop operation to remove and return the top element of the stack
func (s *Stack) Pop() (int, error) {
	if s.top == nil {
		return 0, fmt.Errorf("stack is empty")
	}
	topValue := s.top.data
	s.top = s.top.next
	s.size--
	return topValue, nil
}

// Peek operation to return the top element without removing it
func (s *Stack) Peek() (int, error) {
	if s.top == nil {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.top.data, nil
}

// IsEmpty operation to check if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

// Size operation to return the number of elements in the stack
func (s *Stack) Size() int {
	return s.size
}

func stackLinkedListExample() {
	stack := &Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Println("Stack after pushing elements:")
	temp := stack.top
	for temp != nil {
		fmt.Printf("%d -> ", temp.data)
		temp = temp.next
	}
	fmt.Println("NULL")

	top, _ := stack.Pop()
	fmt.Println("Popped element:", top)

	peek, _ := stack.Peek()
	fmt.Println("Top element (peek):", peek)
	fmt.Println("Is stack empty?", stack.IsEmpty())
	fmt.Println("Stack size:", stack.Size())
}
