package leetcode

import (
	"fmt"
	"sync"
)

// You need to alternate between two goroutines:

// One prints "foo"

// One prints "bar"
// And repeat that sequence n times → "foobarfoobarfoobar...".

type FooBarBasic struct {
	n       int
	fooChan chan struct{}
	barChan chan struct{}
}

func NewFooBarBasic(n int) *FooBarBasic {
	return &FooBarBasic{
		n:       n,
		fooChan: make(chan struct{}, 1),
		barChan: make(chan struct{}, 1),
	}
}

func (fb *FooBarBasic) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.fooChan // Wait until it's foo's turn
		printFoo()
		fb.barChan <- struct{}{} // Signal bar to run
	}
}

func (fb *FooBarBasic) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.barChan // Wait until it's bar's turn
		printBar()
		fb.fooChan <- struct{}{} // Signal foo to run
	}
}
func main() {
	n := 2
	fb := NewFooBarBasic(n)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fb.Foo(func() { fmt.Print("foo") })
	}()

	go func() {
		defer wg.Done()
		fb.Bar(func() { fmt.Print("bar") })
	}()

	// Start the foo/bar loop by signaling foo
	fb.fooChan <- struct{}{}

	wg.Wait()
}

// How It Works:
// We create two channels: fooChan and barChan to coordinate turns.

// Initially, we signal fooChan to start.

// After printing "foo", it signals barChan, which then prints "bar" and signals back to fooChan.

// This continues for n iterations.

// Output (for n := 2):
// foobarfoobar
