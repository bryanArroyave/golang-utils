package customerrors

import "fmt"

type MinLengthError struct {
	value int
}

func NewMinLengthError(value int) *MinLengthError {
	return &MinLengthError{value: value}
}

func (e *MinLengthError) Error() string {
	return fmt.Sprintf("value must have a minimum length of %d", e.value)
}
