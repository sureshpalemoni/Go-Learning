package channels

import "fmt"


func squares(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}
	close(c)
}

func ClosedChannels() {
	fmt.Println("Main goroutine has started")
	c := make(chan int)

	go squares(c) // start goroutine

	for {
		val, ok := <-c
		if ok == false {
			fmt.Println(val, ok, "<-- loop broke!")
			break
		} else {
			fmt.Println(val, ok)
		}
	}

	fmt.Println("main routine stopped")
}