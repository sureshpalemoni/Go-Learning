package errorint

import (
	"fmt"
	"errors"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return (a/b), nil
}

func Test1() {
	if result, err := Divide(4,0); err != nil {
		fmt.Println("Error Occured: ", err)
	} else {
		fmt.Println("4/0 is ", result)	
	}
}