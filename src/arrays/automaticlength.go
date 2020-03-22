package arrays

import "fmt"

func AutomaticArray() {
	greetings := [...]string {
		"Good Morning!",
		"Guten Morgen!",
		"Good evening!",
		"Good Night!",
		"Happy Learning",
	}

	fmt.Println(greetings)
	fmt.Println(len(greetings)) //To find the length of the array
}