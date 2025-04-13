package concepts

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine")
}

func goRoutineMain() {
	go sayHello()           // runs in a new goroutine
	time.Sleep(time.Second) // wait for goroutine to finish
}
