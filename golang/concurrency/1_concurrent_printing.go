package concurrency

import (
	"fmt"
	"sync"
)

// Function to print numbers from start to limit, using a channel to synchronize with other threads
func printNumbers(id, start, limit int, myTurn <-chan bool, nextTurn chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done() // Signal completion once the function finishes

	for i := start; i <= limit; i += 3 {
		<-myTurn // Wait for the signal to proceed (block until it's this goroutine's turn)

		fmt.Printf("Thread %d: %d\n", id, i) // Print the current number

		nextTurn <- true // Signal the next thread to start
	}
}

func main() {
	// Limit up to which we want to print numbers
	limit := 100

	// Channels to coordinate the turn among three threads
	turn1 := make(chan bool, 1) // Start with thread 1
	turn2 := make(chan bool)
	turn3 := make(chan bool)

	// WaitGroup to ensure main waits for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(3) // We have 3 goroutines

	// Launch three goroutines with different starting numbers (1, 2, 3)
	go printNumbers(1, 1, limit, turn1, turn2, &wg)
	go printNumbers(2, 2, limit, turn2, turn3, &wg)
	go printNumbers(3, 3, limit, turn3, turn1, &wg)

	// Initiate the process by signaling the first goroutine to start
	turn1 <- true

	// Wait for all goroutines to complete
	wg.Wait()
}
