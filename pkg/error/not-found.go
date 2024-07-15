package error

type NotFoundError struct {
	SimpleKVError *SimpleKVError
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		SimpleKVError: NewSimpleKVError(message, 404),
	}
}

func (e *NotFoundError) Error() string {
	return e.SimpleKVError.Error()
}
