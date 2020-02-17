package channels

import "fmt"

func greet(c chan string) {
	// fmt.Println("Hello" + <-c + "!" )
	c <- " Suresh"
}

func New() {
	fmt.Println("New goroutine has started, which is main")
	c := make(chan string)

	go greet(c)

	// c <- " Suresh" // At this point main goroutine is blocked untill data in the channel is read
	// Channel operation starts when dats is either
	// read from it read from it
	fmt.Println( "Hello" + <-c + "!")
	fmt.Println("New goroutine stopped")
}