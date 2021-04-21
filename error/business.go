package error

type BusinessError struct {
	Message    string `json:"message"`
	Internal   error  `json:"-"`
	StatusCode int    `json:"statusCode"`
}

type InternalError struct {
	Message    string `json:"message"`
	Internal   error  `json:"-"`
	StatusCode int    `json:"statusCode"`
}

func NewBusinessError(message string, err error, statusCode int) *BusinessError {
	return &BusinessError{
		message,
		nil,
		statusCode,
	}
}

func (e *BusinessError) Error() string {
	return e.Message
}

func NewInternalError(message string, err error, statusCode int) *InternalError {
	return &InternalError{
		message,
		nil,
		statusCode,
	}
}

func (e *InternalError) Error() string {
	return e.Message
}
