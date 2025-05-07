package concurrency

import (
	"fmt"
	"sync"
)

func printNumbersSimpleFunc(start, end int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		fmt.Println(i)
	}
}

func mainExample() {
	var wg sync.WaitGroup

	// Split 1-100 across 3 goroutines
	wg.Add(3)

	go printNumbersSimpleFunc(1, 33, &wg)   // Goroutine 1: 1 to 33
	go printNumbersSimpleFunc(34, 66, &wg)  // Goroutine 2: 34 to 66
	go printNumbersSimpleFunc(67, 100, &wg) // Goroutine 3: 67 to 100

	wg.Wait()
}
