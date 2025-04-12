package container

// Circular Ring (Circular List)
// Method	Description
// ring.New(n)	Creates a ring of size n
// r.Value	Gets/sets value at current node
// r.Next()	Moves to next element
// r.Prev()	Moves to previous element
// r.Do(func(v interface{}))	Iterates over the ring
// r.Link(s)	Connects two rings
// r.Unlink(n)	Removes n elements from the ring, returns ring

import (
	"container/ring"
	"fmt"
)

func ringExample() {
	r := ring.New(3)
	for i := 1; i <= 3; i++ {
		r.Value = i
		r = r.Next()
	}

	r.Do(func(v interface{}) {
		fmt.Println(v)
	})
}
