package main

import (
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
	}

	if cerrObj, ok := cerr.ToCerr(aErr); ok {
		fmt.Println("cerrObj", cerrObj.ToJSON())
	} else {
		fmt.Println("couldn't convert aErr object to Cerr.")
	}
}

func main() {
	CerrHelloWorld()
	CerrTypeCheck()
}
