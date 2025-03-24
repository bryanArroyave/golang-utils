package customerrors

import "fmt"

type MinLengthError struct {
	value int
}

func NewMinLengthError(value int) *MinLengthError {
	return &MinLengthError{value: value}
}

func (e *MinLengthError) Error() string {
	return fmt.Sprintf("min length error: %d", e.value)
}
