package channels

import (
	"fmt"
	"time"
)

func init() {
	start = time.Now()
}

func serviceTimeout1(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello From Service 1"
}

func serviceTimeout2(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello From Service 2"
}

func TimeoutChannel() {
	fmt.Println("main goroutine gas started at ", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go serviceTimeout1(chan1)
	go serviceTimeout2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1 at ", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from service 2 at ", res, time.Since(start))
	case <-time.After(2 * time.Second):
		fmt.Println("No response received, timeout", time.Since(start))		
	}

	fmt.Println("Main goroutine hasa stopped at ", time.Since(start))
}