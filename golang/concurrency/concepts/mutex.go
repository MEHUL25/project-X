package concepts

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func mutexMain() {
	fmt.Println("1️⃣ sync.Mutex: Locking for write")
	mutexExample()

	fmt.Println("\n2️⃣ sync.RWMutex: Multiple readers, one writer")
	rwMutexExample()

	fmt.Println("\n3️⃣ sync/atomic: Lock-free counter increment")
	atomicExample()
}

// //////////////////////////////////////////////////
// 1️⃣ Basic Mutex Example (exclusive lock)
// //////////////////////////////////////////////////
func mutexExample() {
	var mu sync.Mutex
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("🔐 Final counter (mutex):", counter)
}

// //////////////////////////////////////////////////
// 2️⃣ RWMutex Example: Multiple Readers, One Writer
// //////////////////////////////////////////////////
func rwMutexExample() {
	var rwmu sync.RWMutex
	data := make(map[string]string)
	data["version"] = "1.0"

	var wg sync.WaitGroup

	// Writer: changes the value once
	wg.Add(1)
	go func() {
		defer wg.Done()
		rwmu.Lock()
		fmt.Println("✍️ Writer acquiring lock")
		time.Sleep(time.Millisecond * 200)
		data["version"] = "2.0"
		fmt.Println("✅ Writer updated version")
		rwmu.Unlock()
	}()

	// Readers: can read concurrently
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			rwmu.RLock()
			fmt.Printf("👀 Reader %d sees version: %s\n", id, data["version"])
			time.Sleep(time.Millisecond * 100)
			rwmu.RUnlock()
		}(i)
	}

	wg.Wait()
}

// //////////////////////////////////////////////////
// 3️⃣ sync/atomic: Lock-Free Counter (very fast!)
// //////////////////////////////////////////////////
func atomicExample() {
	var counter int64 = 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println("⚛️ Final counter (atomic):", counter)
}
