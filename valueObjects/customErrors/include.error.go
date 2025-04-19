package customerrors

import "fmt"

type IncludeError[T any] struct {
	values []T
}

func NewIncludeError[T any](values []T) *IncludeError[T] {
	return &IncludeError[T]{values: values}
}

func (e *IncludeError[T]) Error() string {
	return fmt.Sprintf("value must be one of: %v", e.values)
}
