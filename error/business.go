package error

type BusinessError struct {
	Message  string `json:"message"`
	Internal error  `json:"-"`
}

func NewBusinessError(message string) *BusinessError {
	return &BusinessError{
		message,
		nil,
	}
}

func (e *BusinessError) SetInternal(err error) *BusinessError {
	e.Internal = err
	return e
}

func (e *BusinessError) Error() string {
	return e.Message
}
