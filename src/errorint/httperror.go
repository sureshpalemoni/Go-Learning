package errorint

import "fmt"

type HttpError struct {
	status int
	method string
}

func (httpError *HttpError) Error() string {
	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.", 
		httpError.method, httpError.status)
}

func getUserEmail(userID int) (string, error) {
	return "", &HttpError{403, "GET"}
}

func HTTPTest() {
	if email, err := getUserEmail(1); err != nil {
		fmt.Println(err)
		if errVal := err.(*HttpError); errVal.status == 403 {
			fmt.Println("Performing session cleaning..")
		}
	} else {
		fmt.Println("User email is", email)
	}
}