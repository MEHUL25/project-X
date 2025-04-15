package leetcode

import (
	"fmt"
	"sync"
)

type FooBar struct {
	n    int
	mu   sync.Mutex
	cond *sync.Cond
	turn bool // true => foo's turn, false => bar's turn
}

func NewFooBar(n int) *FooBar {
	fb := &FooBar{
		n:    n,
		turn: true,
	}
	fb.cond = sync.NewCond(&fb.mu)
	return fb
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		fb.mu.Lock()
		for !fb.turn {
			fb.cond.Wait()
		}
		printFoo()
		fb.turn = false
		fb.cond.Signal()
		fb.mu.Unlock()
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		fb.mu.Lock()
		for fb.turn {
			fb.cond.Wait()
		}
		printBar()
		fb.turn = true
		fb.cond.Signal()
		fb.mu.Unlock()
	}
}
func alternatePrintSyncCond() {
	n := 2
	fb := NewFooBar(n)
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
	wg.Wait()
}
