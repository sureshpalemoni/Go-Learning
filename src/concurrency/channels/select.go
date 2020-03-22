/* Select is just like switch without any
input argument but only used with channels
Select statement is used to perform an operation
on only one channel out of many, condiionally
selected by case block
*/

package channels

import (
	"fmt"
	"time"
)

var start time.Time
func init() {
	start = time.Now()
}

func service1(c chan string) {
	time.Sleep( 3 * time.Second)
	c <- "Hello from the service 1"
}

func service2(c chan string) {
	time.Sleep( 4 * time.Second)
	c <- "Hello from the service 2"
}

func SelectChannel() {
	fmt.Println("Main goroutine started at ", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	// Here only one channel operation will be executed
	// when both the channels unblocked at the same time,
	// random channel is selected
	select {
	case res := <-chan1:
		fmt.Println("Response from the service 1 ", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from the service 2 ", res, time.Since(start))
	}
		
	fmt.Println("Main goroutine has stopped at ", time.Since(start))
}
