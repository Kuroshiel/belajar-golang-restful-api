package exception

// Error Handler Golang RESTful API

// ERROR : NOT FOUND ERROR
type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}
