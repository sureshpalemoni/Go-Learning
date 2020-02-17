package channels

import "fmt"

func LengthCapacity() {
	c := make(chan int, 3)
	c <-1
	c <-2

	fmt.Printf("Length of channel c is %v and capacity of channel c is %v\n", len(c), cap(c))
}