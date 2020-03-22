package channels

import "fmt"


func greetDataType(c chan string) {
	fmt.Println("Hello " + <-c + " !!")
}

func greeter(cc chan chan string) {
	c := make(chan string)
	cc <- c
}

func ChannelDatatype() {
	fmt.Println("Main goroutine started")

	cc := make(chan chan string)

	go greeter(cc)

	c := <-cc

	go greetDataType(c)

	c <- "Suresh"

	fmt.Println("Main Goroutine has stopped")
}