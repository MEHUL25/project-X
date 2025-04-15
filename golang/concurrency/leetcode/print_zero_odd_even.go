package leetcode

import (
	"fmt"
	"sync"
)

// https://leetcode.com/problems/print-zero-even-odd/description/?envType=problem-list-v2&envId=concurrency

type ZeroEvenOdd struct {
	n     int
	curr  int
	state int // 0 = zero's turn, 1 = odd's turn, 2 = even's turn
	mu    sync.Mutex
	cond  *sync.Cond
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	z := &ZeroEvenOdd{
		n:     n,
		curr:  1,
		state: 0,
	}
	z.cond = sync.NewCond(&z.mu)
	return z
}

func (z *ZeroEvenOdd) zero(printNumber func(int)) {
	for i := 0; i < z.n; i++ {
		z.mu.Lock()
		for z.state != 0 {
			z.cond.Wait()
		}
		printNumber(0)
		if z.curr%2 == 1 {
			z.state = 1 // odd's turn
		} else {
			z.state = 2 // even's turn
		}
		z.cond.Signal()
		z.mu.Unlock()
	}
}

func (z *ZeroEvenOdd) even(printNumber func(int)) {
	for {
		z.mu.Lock()
		for z.state != 2 && z.curr <= z.n {
			z.cond.Wait()
		}
		if z.curr > z.n {
			z.mu.Unlock()
			return
		}
		printNumber(z.curr)
		z.curr++
		z.state = 0 // back to zero
		z.cond.Signal()
		z.mu.Unlock()
	}
}

func (z *ZeroEvenOdd) odd(printNumber func(int)) {
	for {
		z.mu.Lock()
		for z.state != 1 && z.curr <= z.n {
			z.cond.Wait()
		}
		if z.curr > z.n {
			z.mu.Unlock()
			return
		}
		printNumber(z.curr)
		z.curr++
		z.state = 0 // back to zero
		z.cond.Signal()
		z.mu.Unlock()
	}
}

func solution1() {
	n := 5
	zeo := NewZeroEvenOdd(n)

	print := func(x int) {
		fmt.Print(x)
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		zeo.zero(print)
	}()
	go func() {
		defer wg.Done()
		zeo.even(print)
	}()
	go func() {
		defer wg.Done()
		zeo.odd(print)
	}()

	wg.Wait()
}
