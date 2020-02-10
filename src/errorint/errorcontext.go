package errorint

import "fmt"

type UnauthorizedErrorContext struct {
	UserId int
	SessionID int
	error
}

func (err *UnauthorizedErrorContext) isLoggedIn() bool {
	return err.SessionID != 0
}

func validateUserContext(UserId int) error {
	err := fmt.Errorf("Session Invalid for user id %d", UserId)
	return &UnauthorizedErrorContext{UserId, 0, err}
}


func ContextError() {
	err := validateUserContext(1)
	if err != nil {
		if errVal, ok := err.(*UnauthorizedErrorContext); ok {
			fmt.Println("Is User logged in: ", errVal.isLoggedIn())
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("user is allowed to perform this action")
	}
}