package channels

import (
	"fmt"
	"time"
)

// var start time.Time .. defined in select.go

func init() {
	start = time.Now()
}

func serviceDef1(c chan string) {
	fmt.Println("ServiceDef1 started", time.Since(start))
	c <- "Hello from the service 1"
}

func serviceDef2(c chan string) {
	fmt.Println("ServiceDef2 started", time.Since(start))
	c <- "Hello from the service 2"
}

func SelectDefChannel() {
	fmt.Println("Main goroutine started at ", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go serviceDef1(chan1)
	go serviceDef2(chan2)


	// time.Sleep( 3 * time.Second)

	select {
	case res := <-chan1:
		fmt.Println("Response from the serviceDef 1 ", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from the serviceDef 2 ", res, time.Since(start))
	default:
		fmt.Println("No response received ", time.Since(start))
	}
		
	fmt.Println("Main goroutine has stopped at ", time.Since(start))
}


