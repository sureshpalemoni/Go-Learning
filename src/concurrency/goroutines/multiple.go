package goroutines

import (
	"fmt"
	"time"
)

var start time.Time 

func init() {
	start = time.Now()
	fmt.Println(start)
}

func getChars(s string) {
	for _, c := range s {
		fmt.Printf("%c at time %v\n", c, time.Since(start))
		time.Sleep(10 * time.Millisecond)
	}
}

func getDigits( s []int) {
	for _, d := range s {
		fmt.Printf("%d", d)
		time.Sleep(30 * time.Millisecond)
	}
}

func Multiple() {
	fmt.Println("Main execution has started at time", time.Since(start))

	go getChars("Hello")

	go getDigits([]int{1,2,3,4,5})

	time.Sleep(200 * time.Millisecond)

	fmt.Println("\n Main executin stopped at time", time.Since(start))
}