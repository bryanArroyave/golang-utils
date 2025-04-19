package customerrors

import "fmt"

type MaxLengthError struct {
	value int
}

func NewMaxLengthError(value int) *MaxLengthError {
	return &MaxLengthError{value: value}
}

func (e *MaxLengthError) Error() string {
	return fmt.Sprintf("max length exceeded: value cannot be longer than %d characters", e.value)
}
