package leetcode

import (
	"fmt"
	"sync"
)

type Foo struct {
	firstDone  chan struct{}
	secondDone chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
		firstDone:  make(chan struct{}),
		secondDone: make(chan struct{}),
	}
}

func (f *Foo) First(printFirst func()) {
	printFirst()
	close(f.firstDone) // Notify that first is done
}

func (f *Foo) Second(printSecond func()) {
	<-f.firstDone // Wait until first is done
	printSecond()
	close(f.secondDone) // Notify that second is done
}

func (f *Foo) Third(printThird func()) {
	<-f.secondDone // Wait until second is done
	printThird()
}

func callerExample() {
	foo := NewFoo()
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		foo.First(func() { fmt.Print("first") })
	}()
	go func() {
		defer wg.Done()
		foo.Second(func() { fmt.Print("second") })
	}()
	go func() {
		defer wg.Done()
		foo.Third(func() { fmt.Print("third") })
	}()

	wg.Wait()
}

// firstDone is a channel used to signal that First() has completed.

// secondDone is used to signal that Second() has completed.

// Third() waits on secondDone, ensuring Second() is completed before continuing.

// Closing a channel in Go acts as a one-time signal that it’s "done".

// Sample Output (Regardless of goroutine scheduling):
// firstsecondthird
