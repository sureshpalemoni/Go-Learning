/* Usual Chanels are bidirectional
which means, we can send data to channel and read from it
Go, provides unidirectional channels for example:
roc := make(<-chan int) - Read only channel
soc := make(chan<- int)	 - recive only channel
You may also convert a bi-directonal channel to unidirectional channel
*/

package channels

import "fmt"

func greetUni(roc <-chan string) { //Converting to unidirectonal
	fmt.Println("Hello " + <-roc + " How are you")
}

func UnidrectionChannel() {
	fmt.Println("Main goroutine started")
	c := make(chan string) // Initialised as bi-directioanl channel

	go greetUni(c)

	c <- "Suresh" 

	fmt.Println("Main goroutine stopped")
}