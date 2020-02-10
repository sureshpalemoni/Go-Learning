package structs

import "fmt"

type Data struct {
	string
	int
	bool
}

func AnonymFields() {
	sample1 := Data{"Monday", 100, true}
	sample1.bool = false

	fmt.Println(sample1)
}