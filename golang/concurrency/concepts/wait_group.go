package concepts

import (
	"fmt"
	"sync"
	"time"
)

func waitGroupMain() {
	fmt.Println("1️⃣ Basic WaitGroup usage")
	basicWaitGroup()

	fmt.Println("\n2️⃣ WaitGroup with anonymous goroutines")
	waitGroupWithAnon()

	fmt.Println("\n3️⃣ WaitGroup + channel to collect results")
	waitGroupWithChannel()

	fmt.Println("\n4️⃣ WaitGroup + iteration trap (closure capture)")
	waitGroupWithLoopFix()

	fmt.Println("\n5️⃣ WaitGroup mistake: Add inside goroutine (⚠️ BAD PRACTICE)")
	waitGroupAddInsideGoroutine()
}

// 1️⃣ Basic: Wait for 3 goroutines to finish
func basicWaitGroup() {
	var wg sync.WaitGroup

	wg.Add(3)
	go worker("A", &wg)
	go worker("B", &wg)
	go worker("C", &wg)

	wg.Wait()
	fmt.Println("✅ All basic workers done")
}

func worker(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 300)
	fmt.Println("👷 Worker", name, "done")
}

// 2️⃣ Anonymous goroutines with WaitGroup
func waitGroupWithAnon() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("🔧 Anon Goroutine %d running\n", id)
			time.Sleep(time.Millisecond * 200)
		}(i)
	}
	wg.Wait()
	fmt.Println("✅ All anonymous goroutines done")
}

// 3️⃣ WaitGroup + Channel for collecting results
func waitGroupWithChannel() {
	var wg sync.WaitGroup
	results := make(chan string, 3)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * 100)
			results <- fmt.Sprintf("📦 Result from task %d", id)
		}(i)
	}

	wg.Wait()
	close(results)

	for res := range results {
		fmt.Println(res)
	}
}

// 4️⃣ Closure trap in loops and fixing it
func waitGroupWithLoopFix() {
	var wg sync.WaitGroup
	names := []string{"alpha", "beta", "gamma"}

	for _, name := range names {
		wg.Add(1)
		go func(n string) { // capture value explicitly
			defer wg.Done()
			time.Sleep(time.Millisecond * 100)
			fmt.Println("🔁 Name:", n)
		}(name)
	}
	wg.Wait()
	fmt.Println("✅ Loop closure issue handled")
}

// 5️⃣ WRONG: Adding inside goroutine – causes race condition (for learning only)
func waitGroupAddInsideGoroutine() {
	var wg sync.WaitGroup
	badExample := func() {
		for i := 0; i < 3; i++ {
			go func(id int) {
				// ⚠️ Wrong: Add inside goroutine (may race)
				wg.Add(1)
				defer wg.Done()
				time.Sleep(time.Millisecond * 100)
				fmt.Println("💥 Bad Goroutine", id)
			}(i)
		}
		wg.Wait()
		fmt.Println("❌ May cause deadlock or panic if raced")
	}
	// Uncomment below only if you understand the risk
	badExample()
	fmt.Println("🚫 Skipped executing badExample() to avoid potential race")
}
