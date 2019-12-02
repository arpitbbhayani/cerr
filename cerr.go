package cerr

// Cerr represents base custom error
type Cerr struct {
	ErrorCode   string `json:"error_code,omitempty"`
	Err         error  `json:"-"`
	Message     string `json:"message,omitempty"`
	IsTimeout   bool   `json:"is_timeout"`
	IsRetryable bool   `json:"is_retryable,omitempty"`
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

// ToCerr tries to convert error into Cerr
// and returns cerr object and an boolean signifying if conversion was possible or not.
func ToCerr(err error) (*Cerr, bool) {
	obj, ok := err.(*Cerr)
	return obj, ok
}
