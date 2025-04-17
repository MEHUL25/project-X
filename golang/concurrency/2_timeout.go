package concurrency

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Simulated long-running task
func longTask(id int, duration time.Duration) string {
	time.Sleep(duration)
	return fmt.Sprintf("Task %d finished", id)
}

// runTask simulates a background task
func runTask(label string, delay time.Duration) <-chan string {
	out := make(chan string)
	go func() {
		time.Sleep(delay)
		out <- fmt.Sprintf("Finished task: %s", label)
	}()
	return out
}

// 1. Timeout using time.After
func timeoutUsingTimeAfter() {
	fmt.Println("\n🕒 Timeout using time.After")

	select {
	case res := <-runTask("TimeAfter", 2*time.Second):
		fmt.Println("✅ Success:", res)
	case <-time.After(1 * time.Second):
		fmt.Println("⛔ Timeout: task took too long")
	}
}

// 2. Timeout using context.WithTimeout
func timeoutUsingContext() {
	fmt.Println("\n🕒 Timeout using context.WithTimeout")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		res := longTask(2, 2*time.Second)
		ch <- res
	}()

	select {
	case <-ctx.Done():
		fmt.Println("⛔ Timeout (context):", ctx.Err())
	case res := <-ch:
		fmt.Println("✅ Success:", res)
	}
}

// 3. Select + time.After manually
func timeoutUsingSelect() {
	fmt.Println("\n🕒 Timeout using select + manual time.After")

	done := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		done <- "Task complete"
	}()

	select {
	case msg := <-done:
		fmt.Println("✅ Success:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("⛔ Timeout reached in select")
	}
}

// 4. HTTP client timeout
func timeoutUsingHTTPClient() {
	fmt.Println("\n🕒 Timeout using HTTP Client")

	client := http.Client{
		Timeout: 1 * time.Second,
	}

	_, err := client.Get("https://httpstat.us/200?sleep=2000") // 2s delay
	if err != nil {
		fmt.Println("⛔ HTTP Timeout/Error:", err)
		return
	}
	fmt.Println("✅ HTTP request succeeded")
}

func timeoutExample() {
	timeoutUsingTimeAfter()
	timeoutUsingContext()
	timeoutUsingSelect()
	timeoutUsingHTTPClient()
}

// Output :

// 🕒 Timeout using time.After
// ⛔ Timeout: task took too long

// 🕒 Timeout using context.WithTimeout
// ⛔ Timeout (context): context deadline exceeded

// 🕒 Timeout using select + manual time.After
// ⛔ Timeout reached in select

// 🕒 Timeout using HTTP Client
// ⛔ HTTP Timeout/Error: Get "...": context deadline exceeded
