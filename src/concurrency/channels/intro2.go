package channels

import "fmt"

// Channels by default are pointers
// You pass channels as an arugument to the function or methods
func Intro2() {
	c := make(chan int) // make will create ready to use channel
	fmt.Printf("type of 'c' is %T\n", c)
	fmt.Printf("value of 'c' is %v\n", c)
}