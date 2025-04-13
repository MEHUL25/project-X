package concepts

import (
	"fmt"
	"sync"
	"time"
)

func unbufferedChannelDemo() {
	fmt.Println("\n🔸 Unbuffered Channel Demo")
	ch := make(chan int)
	go func() {
		ch <- 10
	}()
	val := <-ch
	fmt.Println("Received:", val)
}

func bufferedChannelDemo() {
	fmt.Println("\n🔸 Buffered Channel Demo")
	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "world"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func directionalChannelDemo() {
	fmt.Println("\n🔸 Directional Channel Demo")

	sender := func(ch chan<- int) {
		ch <- 99
	}
	receiver := func(ch <-chan int) {
		fmt.Println("Received:", <-ch)
	}

	ch := make(chan int)
	go sender(ch)
	receiver(ch)
}

func selectDemo() {
	fmt.Println("\n🔸 Select Statement Demo")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "from ch1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}
}

func closeChannelDemo() {
	fmt.Println("\n🔸 Close Channel Demo")

	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Println("Received:", val)
	}
}

func closedChannelReadDemo() {
	fmt.Println("\n🔸 Read from Closed Channel Demo")
	ch := make(chan int, 1)
	ch <- 10
	close(ch)
	fmt.Println(<-ch) // 10
	fmt.Println(<-ch) // 0 (zero value)
}

func sendToClosedChannelDemo() {
	fmt.Println("\n🔸 Panic on Sending to Closed Channel")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	ch := make(chan int)
	close(ch)
	ch <- 1 // PANIC
}

func loopClosureProblemDemo() {
	fmt.Println("\n🔸 Loop Closure Problem Demo (Fixed)")

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("Goroutine:", n)
		}(i)
	}
	wg.Wait()
}

func pipelinePatternDemo() {
	fmt.Println("\n🔸 Pipeline Pattern Demo")

	gen := func(nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			for _, n := range nums {
				out <- n
			}
			close(out)
		}()
		return out
	}

	square := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n * n
			}
			close(out)
		}()
		return out
	}

	for n := range square(gen(2, 3, 4)) {
		fmt.Println("Square:", n)
	}
}

func goChannels() {
	unbufferedChannelDemo()
	bufferedChannelDemo()
	directionalChannelDemo()
	selectDemo()
	closeChannelDemo()
	closedChannelReadDemo()
	sendToClosedChannelDemo()
	loopClosureProblemDemo()
	pipelinePatternDemo()
}
