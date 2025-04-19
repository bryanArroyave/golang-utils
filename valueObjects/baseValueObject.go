package valueobjects

import (
	"errors"

	githuberrors "github.com/pkg/errors"
)

type BaseValueObject[T any] struct {
	name     string
	value    *T
	errors   []error
	optional bool
	validate func()
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

func (s *BaseValueObject[T]) InecureValue() T {
	if s.value == nil {
		var zero T
		return zero
	}
	return *s.value
}

func (s *BaseValueObject[T]) AddError(err error) {
	s.errors = append(s.errors, githuberrors.Wrap(err, s.name))
}
