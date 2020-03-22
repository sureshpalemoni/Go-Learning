package channels

import (
	"fmt"
	"time"
)

func init() {
	start = time.Now()
}

func DefaultWithDeadlock() {
	fmt.Println("Main goroutine has started at ", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	select {
	case res := <-chan1:
		fmt.Println("Response from channel1 ", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from channel2 ", res, time.Since(start))
	default:
		fmt.Println("No gorouines available to send data ", time.Since(start))	
	}

	fmt.Println("Main goroutine stopped ", time.Since(start))
}