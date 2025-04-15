package leetcode

// https://leetcode.com/problems/fizz-buzz-multithreaded/description/?envType=problem-list-v2&envId=concurrency

import (
	"fmt"
	"sync"
)

type FizzBuzz struct {
	n    int
	curr int
	mu   sync.Mutex
}

func NewFizzBuzz(n int) *FizzBuzz {
	return &FizzBuzz{
		n:    n,
		curr: 1,
	}
}

func (fb *FizzBuzz) fizz(printFizz func()) {
	for {
		fb.mu.Lock()
		if fb.curr > fb.n {
			fb.mu.Unlock()
			return
		}
		if fb.curr%3 == 0 && fb.curr%5 != 0 {
			printFizz()
			fb.curr++
		}
		fb.mu.Unlock()
	}
}

func (fb *FizzBuzz) buzz(printBuzz func()) {
	for {
		fb.mu.Lock()
		if fb.curr > fb.n {
			fb.mu.Unlock()
			return
		}
		if fb.curr%5 == 0 && fb.curr%3 != 0 {
			printBuzz()
			fb.curr++
		}
		fb.mu.Unlock()
	}
}

func (fb *FizzBuzz) fizzbuzz(printFizzBuzz func()) {
	for {
		fb.mu.Lock()
		if fb.curr > fb.n {
			fb.mu.Unlock()
			return
		}
		if fb.curr%3 == 0 && fb.curr%5 == 0 {
			printFizzBuzz()
			fb.curr++
		}
		fb.mu.Unlock()
	}
}

func (fb *FizzBuzz) number(printNumber func(int)) {
	for {
		fb.mu.Lock()
		if fb.curr > fb.n {
			fb.mu.Unlock()
			return
		}
		if fb.curr%3 != 0 && fb.curr%5 != 0 {
			printNumber(fb.curr)
			fb.curr++
		}
		fb.mu.Unlock()
	}
}
func solutionMutex1() {
	n := 15
	fb := NewFizzBuzz(n)

	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		fb.fizz(func() { fmt.Print("fizz ") })
	}()

	go func() {
		defer wg.Done()
		fb.buzz(func() { fmt.Print("buzz ") })
	}()

	go func() {
		defer wg.Done()
		fb.fizzbuzz(func() { fmt.Print("fizzbuzz ") })
	}()

	go func() {
		defer wg.Done()
		fb.number(func(x int) { fmt.Printf("%d ", x) })
	}()

	wg.Wait()
}
