// TODO:
//  - Type assertion test
//  - defining custom types extending Cerr
//  - DEfining error types

package test

import (
	"errors"
	"testing"

	"github.com/arpitbbhayani/cerr"
)

func TestNonCerrConversion(t *testing.T) {
	aerr := errors.New("some random error")
	if _, ok := cerr.ToCerr(aerr); ok {
		t.Error("generic error got converted to cerr. this shouldn't have happened.")
	}
}
