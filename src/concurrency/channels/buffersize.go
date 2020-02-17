package channels

import "fmt"

func suqresBuffered(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func BufferedChannel() {
	fmt.Println("main goorutine started")
	c := make(chan int, 3)

	go suqresBuffered(c)

	c <- 1
	c <- 2
	// Until this point channel can hold, since buffer size is 3. 
	// Main goroutine is unblocked and exists the program
	c <- 3 
	c <- 4 // Blocks the main goroutine here, and emptys the buffer as read operation is thirsty

	fmt.Println("Main goroutine stopped")
}