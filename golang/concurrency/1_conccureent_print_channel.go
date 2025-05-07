package concurrency

import (
	"fmt"
	"sync"
)

func workerExample(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range jobs {
		fmt.Printf("Worker %d: %d\n", id, num)
	}
}

func mainExample1() {
	const total = 100
	const numWorkers = 3

	jobs := make(chan int, total)
	var wg sync.WaitGroup

	// Start 3 worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go workerExample(i, jobs, &wg)
	}

	// Send numbers 1 to 100 into the channel
	for i := 1; i <= total; i++ {
		jobs <- i
	}
	close(jobs) // No more jobs to send

	wg.Wait() // Wait for all workers to finish
}
