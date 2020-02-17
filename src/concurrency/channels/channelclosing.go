package channels

import "fmt"


func greetclose(c chan string) {
	<-c // after this line it will switch to main goroutine
	<-c 
}

func ChannelClosing() {
	fmt.Println("Main routine started")

	c := make(chan string)

	go greetclose(c)
	c <- "Suresh"

	close(c)

	c <- "Palemoni" // Channel is already closed in the above line, it will not start channel op
	fmt.Println("Main routine stopped")
}