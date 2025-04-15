package leetcode

import (
	"fmt"
	"sync"
)

type FizzBuzzBasic struct {
	n          int
	fizzCh     chan int
	buzzCh     chan int
	fizzbuzzCh chan int
	numberCh   chan int
	doneCh     chan struct{}
}

func NewFizzBuzzChannel(n int) *FizzBuzzBasic {
	return &FizzBuzzBasic{
		n:          n,
		fizzCh:     make(chan int),
		buzzCh:     make(chan int),
		fizzbuzzCh: make(chan int),
		numberCh:   make(chan int),
		doneCh:     make(chan struct{}),
	}
}

func (fb *FizzBuzzBasic) Run() {
	go func() {
		for i := 1; i <= fb.n; i++ {
			switch {
			case i%15 == 0:
				fb.fizzbuzzCh <- i
			case i%3 == 0:
				fb.fizzCh <- i
			case i%5 == 0:
				fb.buzzCh <- i
			default:
				fb.numberCh <- i
			}
		}
		close(fb.fizzCh)
		close(fb.buzzCh)
		close(fb.fizzbuzzCh)
		close(fb.numberCh)
	}()
}

func (fb *FizzBuzzBasic) fizz(printFizz func()) {
	for range fb.fizzCh {
		printFizz()
	}
	fb.doneCh <- struct{}{}
}

func (fb *FizzBuzzBasic) buzz(printBuzz func()) {
	for range fb.buzzCh {
		printBuzz()
	}
	fb.doneCh <- struct{}{}
}

func (fb *FizzBuzzBasic) fizzbuzz(printFizzBuzz func()) {
	for range fb.fizzbuzzCh {
		printFizzBuzz()
	}
	fb.doneCh <- struct{}{}
}

func (fb *FizzBuzzBasic) number(printNumber func(int)) {
	for x := range fb.numberCh {
		printNumber(x)
	}
	fb.doneCh <- struct{}{}
}

func solutionChannel1() {
	n := 15
	fb := NewFizzBuzzChannel(n)

	var wg sync.WaitGroup
	wg.Add(4)

	fb.Run()

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

	// Wait for all done signals
	for i := 0; i < 4; i++ {
		<-fb.doneCh
	}
	wg.Wait()
}
