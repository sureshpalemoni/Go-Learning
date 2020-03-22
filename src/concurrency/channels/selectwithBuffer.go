package channels

import (
	"fmt"
	"time"
)

// var start time.Time
// We will use the start variable from the other file

func init() {
	start = time.Now()
}

func SelectWithBuffer() {
	fmt.Println("main goroutine has started at ", time.Since(start))
	chan1 := make(chan string, 2)
	chan2 := make(chan string, 2)

	chan1 <- "value 1"
	chan1 <- "value 2"
	chan2 <- "value 1"
	chan2 <- "value 2"

	// Here main goroutine will be unblocked since the buffere is 2
	// and we are sending 2 values only
	// hence select will select randomly any of the channel
	select  {
	case res := <-chan1:
		fmt.Println("Response from channel 1 ", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Respnse from channel 2 ", res, time.Since(start))		
	}

	fmt.Println("Main goroutine has stopped at ", time.Since(start))
}