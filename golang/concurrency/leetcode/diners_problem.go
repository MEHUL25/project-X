package leetcode

// To solve the Dining Philosophers Problem where no philosopher should starve, we need to design a deadlock-free and starvation-free concurrent algorithm.

import (
	"fmt"
	"sync"
)

// DiningPhilosophers defines the structure holding forks and a semaphore
type DiningPhilosophers struct {
	forks     []sync.Mutex
	semaphore chan struct{}
}

// Constructor initializes the DiningPhilosophers object
func Constructor() DiningPhilosophers {
	return DiningPhilosophers{
		forks:     make([]sync.Mutex, 5),
		semaphore: make(chan struct{}, 4), // Only allow 4 philosophers to attempt eating at once
	}
}

// WantsToEat is the method each philosopher calls when they want to eat
func (d *DiningPhilosophers) WantsToEat(
	philosopher int,
	pickLeftFork func(),
	pickRightFork func(),
	eat func(),
	putLeftFork func(),
	putRightFork func(),
) {
	left := philosopher
	right := (philosopher + 1) % 5

	// Acquire semaphore slot (blocks if already 4 are eating)
	d.semaphore <- struct{}{}

	// Lock forks in a consistent order to avoid deadlock
	if philosopher%2 == 0 {
		d.forks[left].Lock()
		d.forks[right].Lock()
	} else {
		d.forks[right].Lock()
		d.forks[left].Lock()
	}

	// Pick up forks
	pickLeftFork()
	pickRightFork()

	// Eat
	eat()

	// Put down forks
	putRightFork()
	putLeftFork()

	// Unlock forks
	d.forks[left].Unlock()
	d.forks[right].Unlock()

	// Release semaphore slot
	<-d.semaphore
}

func dinnerProblem() {
	dp := Constructor()
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		id := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				dp.WantsToEat(
					id,
					func() { fmt.Printf("Philosopher %d picked up left fork\n", id) },
					func() { fmt.Printf("Philosopher %d picked up right fork\n", id) },
					func() { fmt.Printf("Philosopher %d is eating\n", id) },
					func() { fmt.Printf("Philosopher %d put down left fork\n", id) },
					func() { fmt.Printf("Philosopher %d put down right fork\n", id) },
				)
			}
		}()
	}

	wg.Wait()
}
