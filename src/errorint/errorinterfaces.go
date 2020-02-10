package errorint

import "fmt"

type UserSessionState interface {
	error
	isLoggedin() bool
	getSessionId() int
}

type UnauthorizedErrorInt struct {
	UserId int
	SessionId int
}

func (err *UnauthorizedErrorInt) isLoggedin() bool {
	return err.SessionId != 0
}

func (err *UnauthorizedErrorInt) getSessionId() int {
	return err.SessionId
}

func (httpErr *UnauthorizedErrorInt) Error() string {
	return fmt.Sprintf("User with id %v unauthorized to perform this action", httpErr.UserId)
}

func validateUserInt(UserId int) error {
	return &UnauthorizedErrorInt{UserId, 98383}
}

func ErrorWithInterface() {
	err := validateUserInt(1)
	if err != nil {
		fmt.Println(err)
		if errVal, ok := err.(UserSessionState); ok {
			if errVal.isLoggedin() {
				fmt.Printf("cleaning user session id  %v\n", errVal.getSessionId())
			}
		}
	} else {
		fmt.Println("User is  allowed to perform thsi Interface action")
	}
}

