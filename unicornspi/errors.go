package unicornspi

import "fmt"

type Error interface {
	Message() string
	Error() string
	Unwrap() error
}

type baseError struct {
	message string
	cause   error
}

func (e *baseError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s: %v", e.message, e.cause)
	} else {
		return e.message
	}
}

func (e *baseError) Message() string {
	return e.message
}

func (e *baseError) Unwrap() error {
	return e.cause
}

type OpenError struct {
	baseError
}

func newOpenError(message string, cause error) *OpenError {
	return &OpenError{baseError{
		message: message,
		cause:   cause,
	}}
}

type closeError struct {
	baseError
}

func newCloseError(err error) error {
	return &closeError{
		baseError: baseError{message: "error closing device", cause: err},
	}
}
