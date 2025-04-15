package concurrency

import (
	"fmt"
)

func worker(done chan bool) {
	fmt.Println("Working...")
	// Simulate some work
	// time.Sleep(1 * time.Second)
	fmt.Println("Work done!")
	done <- true // Signal completion
}

func basicSynchronizationUsingChannel() {
	done := make(chan bool)

	go worker(done)

	<-done // Wait for the worker to finish
	fmt.Println("Main exits after worker")
}

func example2() {
	count := 0
	mutex := make(chan struct{}, 1) // acts like a lock

	increment := func() {
		mutex <- struct{}{} // acquire lock
		count++
		<-mutex // release lock
	}

	// Run increments concurrently
	for i := 0; i < 100; i++ {
		go increment()
	}

	// Wait (not production safe, just demo)
	// In production, you'd use sync.WaitGroup or more channel signaling
	fmt.Scanln()
	fmt.Println("Final count:", count)
}
