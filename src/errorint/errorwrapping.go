package errorint

import "fmt"

type UnauthorizedError struct {
	UserID int
	OriginalError error
}

func (httpErr *UnauthorizedError) Error() string {
	return fmt.Sprintf("User unauthorized Error: %v", httpErr.OriginalError)
}

func validateUser( userId int) error {
	err := fmt.Errorf("Session invalid for user id %d", userId)
	return &UnauthorizedError{userId, err}
}

func WrapError() {
	if err := validateUser(1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User is allowed to perform this action")
	}
}