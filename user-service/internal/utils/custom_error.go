package utils

import (
	"fmt"
)

type StatusError struct {
	Er     error  // wrapped error
	Status int    // HTTP status code
	Msg    string // Optional user-facing or dev-friendly context message
}

func (se StatusError) Error() string {
	if se.Msg != "" {
		return fmt.Sprintf("%s: %v", se.Msg, se.Er)
	}
	return se.Er.Error()
}

// Optional: allow unwrapping for errors.As / errors.Is
func (se StatusError) Unwrap() error {
	return se.Er
}

// Optional: for direct access to HTTP status in customHandler
func (se StatusError) HTTPStatus() int {
	return se.Status
}

// NewStatusError creates a structured error with optional user/dev message
func NewStatusError(err error, status int, msg ...string) StatusError {
	finalMsg := ""
	if len(msg) > 0 {
		finalMsg = msg[0]
	}
	return StatusError{
		Er:     err,
		Status: status,
		Msg:    finalMsg,
	}
}
