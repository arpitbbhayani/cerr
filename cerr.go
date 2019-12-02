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

type CerrBuilder struct {
	instance *Cerr
}

func NewCerrBuilder() *CerrBuilder {
	return &CerrBuilder{}
}

func (c *CerrBuilder) SetErrorCode(code string) *CerrBuilder {
	c.instance.ErrorCode = code
	return c
}

func (c *CerrBuilder) SetOriginalError(err error) *CerrBuilder {
	c.instance.Err = err
	return c
}

func (c *CerrBuilder) SetMessage(message string) *CerrBuilder {
	c.instance.Message = message
	return c
}

func (c *CerrBuilder) SetIsTimeout(isTimeout bool) *CerrBuilder {
	c.instance.IsTimeout = isTimeout
	return c
}

func (c *CerrBuilder) SetIsRetryable(isRetryable bool) *CerrBuilder {
	c.instance.IsRetryable = isRetryable
	return c
}

func (c *CerrBuilder) Build() *Cerr {
	return c.instance
}
