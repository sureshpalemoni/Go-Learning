package arrays

import "fmt"

func Range() {
	greetings := [...]string {
		"Good Morning!",
		"Guten Morgen!",
		"Good evening!",
		"Good Night!",
		"Happy Learning",
	}

	for i, v := range greetings {
		fmt.Printf("a[%d] is %q\n", i, v)
	}

	// If you are not interested in index
	for _, v := range greetings {
		fmt.Println(v)
	}
}