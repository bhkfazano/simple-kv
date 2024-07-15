package error

import "fmt"

type SimpleKVError struct {
	Message string
	Code    int
}

func NewSimpleKVError(message string, code int) *SimpleKVError {
	return &SimpleKVError{
		Message: message,
		Code:    code,
	}
}

func (e *SimpleKVError) Error() string {
	return fmt.Sprintf("[SimpleKVError: %d] %s", e.Code, e.Message)
}
