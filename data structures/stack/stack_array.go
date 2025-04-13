package stack

import (
	"fmt"
)

type Stack struct {
	elements []int
}

// Push operation to add an element to the stack
func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

// Pop operation to remove and return the top element of the stack
func (s *Stack) Pop() (int, error) {
	if len(s.elements) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top, nil
}

// Peek operation to return the top element without removing it
func (s *Stack) Peek() (int, error) {
	if len(s.elements) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

// IsEmpty operation to check if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size operation to return the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.elements)
}

func stackImplementation() {
	stack := &Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Println("Stack after pushing elements:", stack.elements)

	top, _ := stack.Pop()
	fmt.Println("Popped element:", top)
	fmt.Println("Stack after popping:", stack.elements)

	peek, _ := stack.Peek()
	fmt.Println("Top element (peek):", peek)
	fmt.Println("Is stack empty?", stack.IsEmpty())
	fmt.Println("Stack size:", stack.Size())
}
