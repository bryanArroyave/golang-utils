package valueobjects

import "errors"

type BaseValueObject[T any] struct {
	value    *T
	errors   []error
	optional bool
	validate func()
}

func NewBaseValueObject[T any](value *T) *BaseValueObject[T] {
	return &BaseValueObject[T]{value: value}
}

func (s *BaseValueObject[T]) isValid() bool {
	return len(s.errors) == 0
}

func (s *BaseValueObject[T]) Value() (T, error) {

	s.validate()

	if !s.isValid() {
		var zero T
		return zero, errors.Join(s.errors...)
	}

	if s.value == nil {
		var zero T
		return zero, nil
	}

	return *s.value, nil
}
