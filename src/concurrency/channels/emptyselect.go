package channels

import "fmt"

func service() {
	fmt.Println("Hello from service!!")
}

func EmptySelect() {
	fmt.Println("main goroutine started")

	go service()

	select {}

	fmt.Println("main goroutine stopped")
}