package channels

import (
	"fmt"
	"time"
)

func sqrtWorker(tasks <-chan int, results chan<- int, instance int) {
	for num := range tasks {
		time.Sleep(time.Millisecond)
		fmt.Printf("worker %v sending result by worker %v\n", instance, instance)
		results <- num * num
	}
}

func WorkerPool() {
	fmt.Println("main goroutine has started")

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// spawning 3 worker routines
	for i := 0; i < 3; i++ {
		go sqrtWorker(tasks, results, i)
	}

	// passing 5 tasks
	for i := 0; i<5; i++ {
		tasks <- i * 2 // non-blocking as buffer capacity is 10
	}

	fmt.Println("Main have 5 tasks now")

	// receiving results from all workers
	for i :=0; i<5; i++ {
		result := <-results //blocking because buffer is empty
		fmt.Println("Main reult ", i, ": ", result)
	}

	fmt.Println("main goroutine stopped")
}