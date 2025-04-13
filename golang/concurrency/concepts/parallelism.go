package concepts

import (
	"fmt"
	"runtime"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task", id, "executing on core", runtime.NumCPU())
}

func parallelismExample() {
	runtime.GOMAXPROCS(4) // Use 4 CPU cores
	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go task(i, &wg) // Parallel execution
	}

	wg.Wait() // Wait for all tasks to complete
}
