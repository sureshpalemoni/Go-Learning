package channels

import "fmt"

func Intro() {
	var c chan int
	// This will print <nil>
	//This is not useful because we can't pass or read data from nil channel
	fmt.Println(c)
}