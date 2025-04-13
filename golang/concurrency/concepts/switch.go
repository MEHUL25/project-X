package concepts

import (
	"fmt"
	"time"
)

func switchMain() {
	// ---------------- 1. Multiple channels ----------------
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("🔹 Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("🔹 Received:", msg2)
	}

	// ---------------- 2. Timeout using select ----------------
	timeoutCh := make(chan string)

	go func() {
		time.Sleep(3 * time.Second) // Deliberately longer than timeout
		timeoutCh <- "Late message"
	}()

	select {
	case msg := <-timeoutCh:
		fmt.Println("🔸 Received:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("🔸 Timeout occurred!")
	}

	// ---------------- 3. Non-blocking select using default ----------------
	nonBlockingCh := make(chan string)

	select {
	case msg := <-nonBlockingCh:
		fmt.Println("🔻 Received:", msg)
	default:
		fmt.Println("🔻 No message, not blocking")
	}

	// ---------------- 4. Closed channel behavior ----------------
	closedCh := make(chan int)
	close(closedCh)

	select {
	case val := <-closedCh:
		fmt.Println("✅ Closed channel returned:", val) // Will return 0 immediately
	default:
		fmt.Println("✅ Nothing to read")
	}

	// ---------------- 5. Fan-in: Combine channels ----------------
	fanIn1 := make(chan string)
	fanIn2 := make(chan string)

	go func() {
		fanIn1 <- "A"
	}()
	go func() {
		fanIn2 <- "B"
	}()

	combined := fanIn(fanIn1, fanIn2)

	fmt.Println("🌟 Fan-in received:", <-combined)
	fmt.Println("🌟 Fan-in received:", <-combined)
}

func fanIn(ch1, ch2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for i := 0; i < 2; i++ { // limit for demo
			select {
			case msg := <-ch1:
				out <- "From ch1: " + msg
			case msg := <-ch2:
				out <- "From ch2: " + msg
			}
		}
	}()
	return out
}
