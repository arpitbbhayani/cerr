package cerr

import "encoding/json"

// Cerr represents base custom error
type Cerr struct {
	ErrorCode string `json:"error_code,omitempty"`
	Err       error  `json:"error,omitempty"`
	Message   string `json:"message,omitempty"`
}

// Error returns the error string
// and makes it compatible with error interface.
func (r *Cerr) Error() string {
	return r.Message
}

// OriginalError returns the original error.
func (r *Cerr) OriginalError() error {
	return r.Err
}

// ToJSON marshals the Cerr instance to JSON string
// making it easier to send it as a REST response.
func (r *Cerr) ToJSON() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func ToCerr(err error) (*Cerr, bool) {
	obj, ok := err.(*Cerr)
	return obj, ok
}
