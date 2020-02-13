package goroutines

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("helllo World")
}

func GoroutineIntro() {
	fmt.Println("Main execution started")

	//call function
	// printHello() // Normal Function
	go printHello() // Goroutine function, it runs in the background

	// To block the main gorouitne to execute 
	//printHello goroutine
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Main execution stopped")
}