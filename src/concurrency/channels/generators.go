package channels

import "fmt"


func fib(length int) <- chan int {
	c := make(chan int, length)

	go func() {
		for i, j := 0, 1; i < length; i, j = i+j, i {
			c <- i
		}
		close(c)
	}()

	return c
}


func Generators() {
	for fn := range fib(10) {
		fmt.Println("Current fibonacci number is: ", fn)
	}

}