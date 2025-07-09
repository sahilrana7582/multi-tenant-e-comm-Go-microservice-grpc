package utils

type StatusError struct {
	Er     error
	Status int
}

func (se StatusError) Error() string {
	return se.Er.Error()
}

func NewStatusError(err error, status int) StatusError {
	return StatusError{
		Er:     err,
		Status: status,
	}
}
