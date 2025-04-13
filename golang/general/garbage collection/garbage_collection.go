package garbage

import (
	"fmt"
	"runtime"
)

func garbageCollectionExample() {
	// Display initial memory stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Initial GC - Alloc:", m.Alloc, "Heap Objects:", m.HeapObjects)

	// Create objects dynamically
	for i := 0; i < 100000; i++ {
		_ = make([]byte, 1024) // Allocate memory
	}

	runtime.ReadMemStats(&m)
	fmt.Println("Before GC - Alloc:", m.Alloc, "Heap Objects:", m.HeapObjects)

	// Manually trigger garbage collection
	runtime.GC()

	// Display memory stats after GC
	runtime.ReadMemStats(&m)
	fmt.Println("After GC - Alloc:", m.Alloc, "Heap Objects:", m.HeapObjects)
}
