package channels

import (
	"fmt"
	"sync"
	"time"
)

func serviceWaitGroup(wg *sync.WaitGroup, instance int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Service called on instance ", instance )
	wg.Done()
}

func WaitGroup() {
	fmt.Println("Main goroutine started")
	var wg sync.WaitGroup // creates waitgroup - empty struct

	// Spawns three goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go serviceWaitGroup(&wg, i)
	}

	wg.Wait()
	fmt.Println("main goroutine has stopped")
}