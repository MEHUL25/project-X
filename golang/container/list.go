package container

// Method	Description
// list.New()	Creates a new doubly linked list
// l.PushFront(value)	Inserts value at the front
// l.PushBack(value)	Inserts value at the back
// l.InsertBefore(value, element)	Inserts value before a given element
// l.InsertAfter(value, element)	Inserts value after a given element
// l.Remove(element)	Removes and returns the element
// l.Front()	Returns the first element
// l.Back()	Returns the last element
// l.Len()	Returns the number of elements
// element.Value	Access the actual value from a list element
// element.Next() / element.Prev()	Traverse next/previous elements

import (
	"container/list"
	"fmt"
)

func listExample() {
	l := list.New()
	l.PushBack("Go")
	l.PushFront("Hello")
	fmt.Println(l.Front().Value) // Hello
}
