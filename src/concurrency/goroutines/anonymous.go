package goroutines

import (
	"fmt"
	"time"
)

func Anonymous() {
	fmt.Println("main exection started")

	// Anonymous goroutine
	go func() {
		fmt.Println("Hello world")
	}()

	// sleep to schedule another goroutine
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Main execution stopped")
}