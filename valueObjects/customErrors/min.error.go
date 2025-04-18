package customerrors

import "fmt"

type MinError[T any] struct {
	value T
}

func NewMinError[T any](value T) *MinError[T] {
	return &MinError[T]{value: value}
}

func (e *MinError[T]) Error() string {
	return fmt.Sprintf("value is below the minimum allowed(%v)", e.value)
}
