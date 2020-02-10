package errorint

import (
	"fmt"
	"errors"
)

type MyError struct {}

func (myErr *MyError) Error() string {
	return "Something bad happened"
}

func Error() {
	myErr := &MyError{}
	fmt.Println(myErr)

	NewError := errors.New("Something might have happened, NEW")
	fmt.Printf("Error is %#v and is of type %T\n", NewError, NewError)
}

