package channels

import (
	"fmt"
	"runtime"
)


func sender(c chan int) {
	c <- 1 // len 1, cap 3
	c <- 2 // len 2, cap 3
	c <- 3 // len 3, cap 3
	c <- 4 // main goroutine blocks here and starts this routine
	close(c)
}

func LengthCapacity2() {
	c := make(chan int, 3)

    fmt.Println("Active Goroutines", runtime.NumGoroutine())
	go sender(c)

	fmt.Printf("Length of channel c is %v and capacity of channel c is %v\n", len(c), cap(c))

	fmt.Println("Active Goroutines", runtime.NumGoroutine())
	for val := range c {
		fmt.Printf("Length of channel c after value %v read is %v\n", val, len(c))
	}
}