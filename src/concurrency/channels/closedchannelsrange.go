package channels

import "fmt"

func squaresWithRange(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}
	close(c)
}

func ClosedChannelsRange() {
	fmt.Println("Main goroutine has started")
	c := make(chan int)

	go squaresWithRange(c) // start goroutine

	// Executes for loop this untill channel closes
	for val := range c {
		fmt.Println(val)
	}

	fmt.Println("main routine stopped")
}