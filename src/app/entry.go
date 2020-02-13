package main

import (
	"concurrency/goroutines"
)

var integers [10]int

func init() {
	for i :=0; i<10; i++ {
		integers[i] = i
	}
}

func main() {
	goroutines.Example2()
}
