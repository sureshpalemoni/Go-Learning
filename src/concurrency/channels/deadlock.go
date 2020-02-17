package channels

import "fmt"

func Deadlock() {
	fmt.Println("main goroutine started")

	c := make(chan string) 
	c <- "Suresh" // Results in a deadlock as there are no goroutines to read the data from the channel

	fmt.Println("Main goroutine stopped")
}