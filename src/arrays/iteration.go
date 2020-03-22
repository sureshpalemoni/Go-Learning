package arrays

import "fmt"

func Iteration() {
	greetings := [...]string {
		"Good Morning!",
		"Guten Morgen!",
		"Good evening!",
		"Good Night!",
		"Happy Learning",
	}

	for i := 0; i < len(greetings); i++ {
		fmt.Printf("a[%d] is %q\n", i, greetings[i])
	}
}