package channels

import "fmt"


func AnonymousGoroutineChannel() {
	fmt.Println("Main Goroutine has started")
	c := make(chan string)
	go func(c chan string) {
		fmt.Println("Hello " + <-c + " !!")
	}(c)

	c <- "Suresh"
	fmt.Println("Main Goroutine has stopped")
}