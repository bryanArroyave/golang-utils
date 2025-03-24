package customerrors

import "fmt"

type MaxError[T any] struct {
	value T
}

func NewMaxError[T any](value T) *MaxError[T] {
	return &MaxError[T]{value: value}
}

func (e *MaxError[T]) Error() string {
	return fmt.Sprintf("max error: %v", e.value)
}
