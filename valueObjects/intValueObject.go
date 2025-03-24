package valueobjects

import (
	customerrors "github.com/bryanArroyave/golang-utils/valueObjects/customErrors"
)

type IntValueObject struct {
	*BaseValueObject[int]
	minValue         int
	hasMinValidation bool
	maxValue         int
	hasMaxValidation bool
}

func NewIntValueObject(value int) *IntValueObject {
	v := &IntValueObject{}

	base := &BaseValueObject[int]{
		value:    &value,
		validate: v.validate,
	}

	v.BaseValueObject = base

	return v
}

func (s *IntValueObject) Max(value int) *IntValueObject {
	s.maxValue = value
	s.hasMaxValidation = true
	return s
}

func (s *IntValueObject) Min(value int) *IntValueObject {

	s.minValue = value
	s.hasMinValidation = true
	return s
}

func (s *IntValueObject) Optional() *IntValueObject {
	s.optional = true
	return s
}

func (s *IntValueObject) validate() {
	s.validateMax()
	s.validateMin()
}

func (s *IntValueObject) validateMin() {

	if *s.value == 0 && s.optional {
		return
	}
	if s.hasMinValidation {
		if *s.value < s.minValue {
			s.errors = append(s.errors, customerrors.NewMinError(s.minValue))
		}
	}
}

func (s *IntValueObject) validateMax() {

	if *s.value == 0 && s.optional {
		return
	}
	if s.hasMaxValidation {
		if *s.value > s.maxValue {
			s.errors = append(s.errors, customerrors.NewMaxError(s.maxValue))
		}
	}
}
