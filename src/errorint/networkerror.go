package errorint

import "fmt"

type NetworkError struct {}

func (e *NetworkError) Error() string {
	return "A network COnnection ABorted"
}

type FileSaveFailedError struct {}

func (e *FileSaveFailedError) Error() string {
	return "The requested file could not be saved"
}

func saveFileToRemote() error {
	result := 2

	if result == 1 {
		return &NetworkError{}
	} else if result == 2 {
		return &FileSaveFailedError{}
	}
	return nil
}

func NetFileError() {
	switch err := saveFileToRemote(); err.(type) {
		case nil:
			fmt.Println("File saved successfully")
		case *NetworkError:
			fmt.Println("Network Error: ", err)
		case *FileSaveFailedError:
			fmt.Println("File Saving Error", err)
	}
} 