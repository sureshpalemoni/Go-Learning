package main

import (
	"concurrency/channels"
)

var integers [10]int

func init() {
	for i :=0; i<10; i++ {
		integers[i] = i
	}
}

func main() {
	channels.UnidrectionChannel()
}
