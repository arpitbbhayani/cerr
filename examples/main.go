package main

import (
	"encoding/json"
	"fmt"

	"github.com/arpitbbhayani/cerr"
)

// CerrHelloWorld shows what we return is in fact an error
func CerrHelloWorld() error {
	return &cerr.Cerr{}
}

// CerrTypeCheck shows how to convert (in possible) error into Cerr
func CerrTypeCheck() {
	errorCode := "ERROR_1"
	var aErr error
	aErr = &cerr.Cerr{
		ErrorCode: errorCode,
		Message:   "some error1 happened",
	}

	if cerrObj, ok := cerr.ToCerr(aErr); ok {
		fmt.Println("cerrObj", cerrObj)
	} else {
		fmt.Println("couldn't convert aErr object to Cerr.")
	}
}

// CerrJSONString shows how to convert Cerr into JSON string
func CerrJSONString() {
	errorCode := "ERROR_1"
	err := &cerr.Cerr{
		ErrorCode: errorCode,
		Err:       fmt.Errorf("unique constraint violation"),
	}
	errJSONString, _ := json.Marshal(err)
	fmt.Println(string(errJSONString))
}

// CerrExtending shows how to create sub-errors of Cerr and check for instance.
func CerrExtending() {
	type UnauthorizedError struct {
		cerr.Cerr
		HTTPCode int `json:"http_code"`
	}

	type PathError struct {
		cerr.Cerr
		Path string `json:"path"`
	}

	err1 := &UnauthorizedError{cerr.Cerr{ErrorCode: "unauthorized", Message: "Unauthorized access forbidden"}, 403}
	err2 := &PathError{cerr.Cerr{ErrorCode: "path_error", Message: "Path is invalid"}, "/home/sbin/does-not-exist"}

	var err error

	err = err1
	switch e := err.(type) {
	case *UnauthorizedError:
		fmt.Println("err1 is of type UnauthorizedError", e)
	case *PathError:
		fmt.Println("err1 is of type PathError", e)
	default:
		fmt.Println("err1 is generic error", e)
	}

	err = err2
	switch e := err.(type) {
	case *UnauthorizedError:
		fmt.Println("err2 is of type UnauthorizedError", e)
	case *PathError:
		fmt.Println("err2 is of type PathError", e)
	default:
		fmt.Println("err2 is generic error", e)
	}
}

func main() {
	CerrHelloWorld()
	CerrTypeCheck()
	CerrJSONString()
	CerrExtending()
}
