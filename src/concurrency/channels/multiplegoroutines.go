package channels

import "fmt"

func square(c chan int) {
	fmt.Println("Reading from Square channel")
	num := <- c
	c <- num * num
}

func cube(c chan int) {
	fmt.Println("Reading from cube channel")
	num := <- c
	c <- num * num * num
}

func MultiGoroutines() {
	fmt.Println("main Goroutine has started")

	suqareChan := make(chan int)
	cubeChan := make(chan int)

	go square(suqareChan)
	go cube(cubeChan)

	testNum := 3

	fmt.Println("Suqare Goroutine has started")
	suqareChan <- testNum

	fmt.Println("Cube Goroutine has started")
	cubeChan <- testNum

	squareVal, cubeVal := <-suqareChan, <-cubeChan
	sum := squareVal + cubeVal

	fmt.Printf("Sum of square and cube of the number %v is %v\n", testNum, sum)
	fmt.Println("main Goroutine has stopped")
}	